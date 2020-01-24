package app

import (
	"github.com/SophieDeBenedetto/golang-microservices/src/api/logger/option_b"
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
	option_b.Info("Starting App!", "step1:starting", "status:pending")
	// logger.Logger.Info("Starting App!", "step1:starting", "status:pending")
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
