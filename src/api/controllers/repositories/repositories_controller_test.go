package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/clients/restclient"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/repositories"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/utils/errors"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

// Mock rest api call
func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

// T1
func TestCreateRepoInvalidJSONRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response) // get mocked context based on this request and response

	// test create repo
	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	// fmt.Println(response.Body.String())

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err) // not expecting an error
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

// T2
func TestCreateRepoErrorFromGithub(t *testing.T) {
	// create mock
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)),
		},
	})

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "some-cool-repo"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response) // get mocked context based on this request and response

	// test create repo
	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)

	// fmt.Println(response.Body.String())

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(t, err) // not expecting an error
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires authentication", apiErr.Message())
}

// T3
func TestCreateRepoNoError(t *testing.T) {
	// create mock
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "some-cool-repo"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response) // get mocked context based on this request and response

	// test create repo
	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	// no error. valid CreateRepo response
	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err) // not expecting an error
	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
