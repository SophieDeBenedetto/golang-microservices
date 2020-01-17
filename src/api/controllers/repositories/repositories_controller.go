package repositories

import (
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/repositories"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/services"
	"github.com/SophieDeBenedetto/golang-microservices/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateRepo enacts the create repo flow given the incoming web request
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := (c.ShouldBindJSON)(&request); err != nil {
		apiError := errors.BadRequestError("invalid JSON body")
		c.JSON(apiError.GetStatus(), apiError)
		return
	}
	result, err := services.RepoService.CreateRepo(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// CreateRepos enacts the create repo flow given the incoming web request
func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := (c.ShouldBindJSON)(&request); err != nil {
		apiError := errors.BadRequestError("invalid JSON body")
		c.JSON(apiError.GetStatus(), apiError)
		return
	}
	result, err := services.RepoService.CreateRepos(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(result.StatusCode, result)
}
