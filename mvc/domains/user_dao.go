package domains

import (
	"fmt"
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {
			ID:        123,
			FirstName: "Sophie",
			LastName:  "DeBenedetto",
			Email:     "sophie@email.com",
		},
	}
)

// GetUser gets the user with the given ID or returns an error
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %d not found", userID),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
