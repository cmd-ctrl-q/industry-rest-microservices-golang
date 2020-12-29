package test_utils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetMockedContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response) // create context based on this response
	c.Request = request                     // assign request to context
	return c
}
