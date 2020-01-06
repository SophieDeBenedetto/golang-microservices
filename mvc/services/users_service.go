package services

import (
	"github.com/SophieDeBenedetto/golang-microservices/mvc/domains"
	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
)

type usersService struct {
}

var (
	// UsersService instance
	UsersService usersService
)

// GetUser gets a user with the given ID
func (u *usersService) GetUser(userID int64) (*domains.User, *utils.ApplicationError) {
	return domains.UserDao.GetUser(userID)
}
