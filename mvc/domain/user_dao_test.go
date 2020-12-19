package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "Not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID, "expecting %v, got %v", 123, user.ID)
	assert.EqualValues(t, "John", user.FirstName, "expecting %v, got %v", "John", user.FirstName)
	assert.EqualValues(t, "Wick", user.LastName, "expecting %v, got %v", "Wick", user.LastName)
	assert.EqualValues(t, "wicker@gmail.com", user.Email, "expecting %v, got %v", "wicker@gmail.com", user.Email)
}
