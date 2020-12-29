package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock) // map[method_url]*Mock
)

type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Err        error
}

// getMockID is used as the key
func getMockID(httpMethod string, url string) string {
	// return fmt.Sprintf("%s_%s", m.HTTPMethod, m.URL)
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMockups() {
	enableMocks = true
}

func FlushMocks() {
	mocks = make(map[string]*Mock)
}

func StopMockups() {
	enableMocks = false
}

func AddMockup(mock Mock) {
	mocks[getMockID(mock.HTTPMethod, mock.URL)] = &mock
}

// Generic HTTP Client - basic post api call.
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	if enableMocks {
		mock := mocks[getMockID(http.MethodPost, url)] // mocking url
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}
		return mock.Response, mock.Err
		// TODO: return local mock without calling any external resources

	}

	jsonBytes, err := json.Marshal(body)
	if err != nil { // if err then the body type interface was not a valid type to marshal to json
		return nil, err // if error then no way of using this body
	}

	// send request against this client
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
