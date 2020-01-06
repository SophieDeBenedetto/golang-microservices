package domains

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

var (
	users = map[int64]*User{
		123: {
			ID:        123,
			FirstName: "Sophie",
			LastName:  "DeBenedetto",
			Email:     "sophie@email.com",
		},
	}
	// UserDao instance
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

// GetUser gets the user with the given ID or returns an error
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("Querying database...")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %d not found", userID),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
