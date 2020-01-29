package app

import (
	logger "github.com/SophieDeBenedetto/golang-microservices/src/api/logger/option_b"
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
	logger.Info("Starting App!", "step1:starting", "status:pending")
	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}
}
