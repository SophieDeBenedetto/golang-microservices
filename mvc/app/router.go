package app

import (
	"github.com/SophieDeBenedetto/golang-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
