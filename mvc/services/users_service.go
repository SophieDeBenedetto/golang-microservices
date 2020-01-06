package services

import (
	"github.com/SophieDeBenedetto/golang-microservices/mvc/domains"
	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domains.User, *utils.ApplicationError) {
	return domains.GetUser(userId)
}
