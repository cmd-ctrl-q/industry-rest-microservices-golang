package services

import (
	"net/http"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/domain"
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

// GetItem ...
func (is *itemsService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {

	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
