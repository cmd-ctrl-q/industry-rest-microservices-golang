package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/clients/restclient"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	fmt.Println(response)
	fmt.Println(err)
	// T. Rest Client Error
	if err != nil {
		// only way to have error here is if we lost connectivity, (e.g.loss of internet conn)
		// did not hit api
		log.Printf("error when trying to create new repo in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	// T. Invalid Response Body
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	defer response.Body.Close() // send response.Body to defer stack

	// if response > 299, then the response is an error
	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		// T. Invalid Error Interface
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				// github api response may have changed.
				Message: "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode // change the status code error
		// Status Code > 299 Error
		return nil, &errResponse
	}

	// response was successful
	var result github.CreateRepoResponse
	// T. Invalid Success Response
	// attempt to use successful response to create a repo
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying to unmarshal github create repo response",
		}
	}

	return &result, nil
}
