package githubprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	accessToken := "ABC123"
	assert.EqualValues(t, "token ABC123", getAuthorizationHeader(accessToken))
}
