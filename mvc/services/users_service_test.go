package services

import (
	"net/http"
	"testing"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/domain"
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {

	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 does not exist",
		}
	}

	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)   // if user is nil, test failed
	assert.NotNil(t, err) // if err is not nil, then there is an error and testfailed
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123,
		}, nil
	}
	user, err := UsersService.GetUser(123)

	assert.Nil(t, err)     // if err is nil, test failed
	assert.NotNil(t, user) // if user is not nil, then there is an error and test failed
	assert.EqualValues(t, 123, user.ID)
}
