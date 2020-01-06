package services

import (
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/domains"
	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
)

type itemsService struct {
}

var (
	// ItemsService instance
	ItemsService itemsService
)

// GetItem gets an item
func (i *itemsService) GetItem(itemID int64) (*domains.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "implement me",
		Status:  http.StatusInternalServerError,
	}
}
