package app

import (
	"github.com/SophieDeBenedetto/golang-microservices/src/api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// StartApp runs the webserver
func StartApp() {
	mapUrls()
	logger.Logger.Info("Starting App!")
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
