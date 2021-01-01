package services

import (
	"net/http"
	"sync"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/config"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/github"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/repositories"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/providers/github_provider"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError)
}

var (
	// RepositoryService allows to mock service interface
	RepositoryService reposServiceInterface
)

// called a single time where you have an import for this package
func init() {
	RepositoryService = &reposService{}
}

// public interface for the service
// func (s *reposService) CreateRepo(request interface{}) (interface{}, error) {}
func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {

	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	// T2
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	// T3
	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}

func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError) {

	// create channels
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	// wait group. control mechanism to block execution until some work has been done
	var wg sync.WaitGroup

	// dispatch and create a control go routine that will handle the results that
	// all the threads created by `for` loop are going to be sending.
	go s.handleRepoResults(&wg, input, output) // listens for events inside input channel

	// 3 requests to process
	for _, current := range requests {
		wg.Add(1) // add 1 slot to the wait group
		// for each request, create a new go routine for it.
		// once we have a function call in a different go routine,
		// the function must not return anything, it send result to a given channel
		// in this case `input` channel
		go s.createRepoConcurrent(current, input) // go routine / channel locked
		// this go routine simply sends data into the `input` channel
	}
	// execution will freeze here until wait group reaches zero.
	// someone needs to notify the wait group every time a job has been completed. see wg.Done()
	wg.Wait() // when all is done ,release lock and close input channel.
	// which will close the `incomingEvent := range input` loop and return the result
	// when wg reaches zero, then execution releases.
	close(input) // when wg = 0, close input channel

	// blocks until some result comes unto the output channel.
	// the only way for something to go into the output channel is from the handleRepoResult() method.
	// when the results go into the output channel. ie output <- results.
	// but the handleRepoResult() must finish processing all the requests to github.
	result := <-output // go routine / channel locks. this says "Hey im going to just wait until i receive something from the output channel"
	// output will lock here until someone writes something into the channel.

	// partial content response
	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil // return result
	// this result and the result variable below in handeRepoResults() are the same.
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	// results that will be sent to a channel
	var results repositories.CreateReposResponse
	// converge all concurrent results into a single one `results` variable

	// for every result coming in from github, append it to the final `results` variable above.
	// iterate incoming results from the input channel
	// once go routine closes, this range finishes because wg = 0.
	for incomingEvent := range input {

		// process results
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)

		// subtract 1 unit from the wait group
		wg.Done() // notifies the wait group when job has completed.

	}
	// once input channel is closed, send `results` variable to output.
	output <- results // sends final result into result var at line 93 result:= <-output
}

// n requests to process, n go routines created.
// increment counter on wait group for every new go routine
func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	// error in request repo name
	// validate error, if err, send error to output channel
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	result, err := s.CreateRepo(input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}
