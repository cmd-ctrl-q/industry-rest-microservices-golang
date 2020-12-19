package services

import (
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/domain"
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

// GetUser calls the UserDao GetUser() in domain/user_dao.go
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {

	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
