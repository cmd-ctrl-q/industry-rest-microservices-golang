package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/clients/restclient"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/repositories"
)

// before running any test cases, we must initialize the mocks for the restclient.
func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

// T1
func TestCreateRepoInvalidInputName(t *testing.T) {
	// no need to flush here bc not using provider
	request := repositories.CreateRepoRequest{} // in domain dir

	result, err := RepositoryService.CreateRepo(request)

	// result should be nil, since were not sending a valid name
	assert.Nil(t, result) // assert the result is nil
	assert.NotNil(t, err) // assert the err is not nil
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

// T2 - when we have an invalid response from github
func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMocks() // using github provider so must flush mock

	// when you have a post req to this url, with error, return this mock with unauthorized and this body
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)),
		},
	})
	// by adding a Name, we skip the first nil return
	request := repositories.CreateRepoRequest{Name: "some-cool-repo"}
	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result) // result should be nil
	assert.NotNil(t, err) // error should not be nil
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMocks() // using github provider so must flush mock

	// when you have a post req to this url, with error, return this mock with unauthorized and this body
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated, // skips 2nd error
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "some-cool-repo", "owner": {"login":"cmd-ctrl-q"}}`)),
		},
	})
	// by adding a Name, we skip the first nil return
	request := repositories.CreateRepoRequest{Name: "some-cool-repo"}
	result, err := RepositoryService.CreateRepo(request)

	// successful return
	assert.NotNil(t, result) // result should not be nil
	assert.Nil(t, err)       // error should be nil
	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "some-cool-repo", result.Name)
	assert.EqualValues(t, "cmd-ctrl-q", result.Owner)
}
