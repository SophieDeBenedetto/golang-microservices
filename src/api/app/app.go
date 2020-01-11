package app

import (
	"fmt"

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
	fmt.Println("Starting app...")
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
