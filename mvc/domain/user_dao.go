package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
)

// database mock
var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "John", LastName: "Wick", Email: "wicker@gmail.com"},
	}

	// UserDao mock
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{} // set UserDao to a userDao struct on initialization
}

type usersDaoInterface interface {
	// GetUser...
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

// GetUser ...
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {

	log.Println("we're accessing the database")

	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		StatusCode: http.StatusNotFound,
		Message:    fmt.Sprintf("user %v does not exist", userID),
		Code:       "not_found",
	}
}
