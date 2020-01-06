package domains

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(1234)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, fmt.Sprintf("User %v not found", 1234), err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, int64(123), user.ID)
	assert.EqualValues(t, "Sophie", user.FirstName)
	assert.EqualValues(t, "DeBenedetto", user.LastName)
	assert.EqualValues(t, "sophie@email.com", user.Email)
}
