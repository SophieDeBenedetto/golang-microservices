package app

import (
	"fmt"
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/controllers"
)

func StartApp() {
	fmt.Println("Starting app...")
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
