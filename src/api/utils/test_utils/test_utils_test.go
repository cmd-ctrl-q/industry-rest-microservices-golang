package test_utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockedContext(t *testing.T) {
	// request := httptest.NewRequest(http.MethodGet, "http://localhost:123/somsething", nil)
	request, err := http.NewRequest(http.MethodGet, "http://localhost:123/somsething", nil)
	assert.Nil(t, err) // assert error must be nil
	response := httptest.NewRecorder()
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockedContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/somsething", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("x-mock")) // pass
	assert.EqualValues(t, "true", c.GetHeader("X-Mock")) // pass
}
