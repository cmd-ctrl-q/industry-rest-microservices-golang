package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test GetUser() when there is no user found
func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	// calls only if user doesn't exist.
	assert.Nil(t, user, "Not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exist", err.Message)

	// if user exists,
	// if user != nil {
	// 	t.Errorf("User exists with id %d", user.ID)
	// }

	// // if no error, ie if user exists
	// if err == nil {
	// 	t.Errorf("user with id %d exist", user.ID)
	// }

	// // if 404
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Errorf("expecting %v, got %v", 404, err.StatusCode)
	// }

}

// check return user that exists
func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123) // user 123 Does exist

	assert.Nil(t, err)     // check if err is nil, if so, user exists
	assert.NotNil(t, user) // check if user is not nil, if so, user exists

	assert.EqualValues(t, 123, user.ID, "expecting %v, got %v", 123, user.ID)
	assert.EqualValues(t, "John", user.FirstName, "expecting %v, got %v", "John", user.FirstName)
	assert.EqualValues(t, "Wick", user.LastName, "expecting %v, got %v", "Wick", user.LastName)
	assert.EqualValues(t, "wicker@gmail.com", user.Email, "expecting %v, got %v", "wicker@gmail.com", user.Email)
}
