package services

import (
	"net/http"
	"testing"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/domains"
	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"

	"github.com/stretchr/testify/assert"
)

type userDaoMock struct {
}

var (
	UserDaoMock     userDaoMock
	getUserFunction func(userID int64) (*domains.User, *utils.ApplicationError)
)

func (m *userDaoMock) GetUser(userID int64) (*domains.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func init() {
	domains.UserDao = &userDaoMock{}
}

func TestGetUserNotFound(t *testing.T) {
	getUserFunction = func(userID int64) (*domains.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Status: 404,
		}
	}
	user, err := UsersService.GetUser(1234)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
}

func TestGetUser(t *testing.T) {
	getUserFunction = func(userID int64) (*domains.User, *utils.ApplicationError) {
		return &domains.User{ID: 123}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, int64(123), user.ID)
}
