package domains

import (
	"fmt"
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {
			Id:        1,
			FirstName: "Sophie",
			LastName:  "DeBenedetto",
			Email:     "sophie@email.como",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %d not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
