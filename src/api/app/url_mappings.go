package app

import "github.com/SophieDeBenedetto/golang-microservices/src/api/controllers/repositories"

func mapUrls() {
	router.POST("/repos", repositories.CreateRepo)
}
