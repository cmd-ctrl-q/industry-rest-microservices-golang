package services

import (
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/domain"
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
)

// GetUser calls the domain GetUser 
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {

	return domain.GetUser(userID)
}
