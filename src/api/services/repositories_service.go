package services

import (
	"strings"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/config"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/github"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/repositories"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/providers/github_provider"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
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
	input.Name = strings.TrimSpace(input.Name)
	// T1
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
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
