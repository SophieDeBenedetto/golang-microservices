package app

import (
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/controllers/examples"
	"github.com/SophieDeBenedetto/golang-microservices/oauth_api/controllers/oauth"
)

func mapUrls() {
	router.GET("/", examples.Index)
	router.POST("/oauth/token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
