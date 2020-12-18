package domain

import (
	"fmt"
	"net/http"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
)

// mock database
var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "John", LastName: "Wick", Email: "wicker@gmail.com"},
	}
)

// GetUser domain is making the query against the database
func GetUser(userID int64) (*User, *utils.ApplicationError) {

	// if no error, return the error back to the serveres.GetUser() func
	if user := users[userID]; user != nil {
		return user, nil
	}
	// return User{}, fmt.Errorf("user %v was not found", userID) // if returning User, must return empty User{}
	// return nil, fmt.Errorf("user %v was not found", userID) // if returning *User, can use nil

	// if error, return the error back to the serveres.GetUser() func
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found.", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
