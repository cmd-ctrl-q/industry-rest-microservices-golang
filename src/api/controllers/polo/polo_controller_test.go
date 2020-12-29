package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/utils/test_utils"
)

// test the constants
func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	Marco(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
