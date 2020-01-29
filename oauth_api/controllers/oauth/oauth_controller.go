package oauth

import (
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/domains/oauth"
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/services"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateAccessToken responds to a request to create a new access token
func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.BadRequestError("invalid JSON body")
		c.JSON(apiError.GetStatus(), apiError)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, token)

}

// GetAccessToken responds to a request to get an access token
func GetAccessToken(c *gin.Context) {
	tokenID := c.Param("token_id")
	token, err := services.OauthService.GetAccessToken(tokenID)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}
	c.JSON(http.StatusOK, token)

}
