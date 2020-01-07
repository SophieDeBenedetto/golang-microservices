package controllers

import (
	"net/http"
	"strconv"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/services"
	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
)

// GetUser gets the requested user
func GetUser(c *gin.Context) {
	// query params: c.Query("param_name")
	userID, err := (strconv.ParseInt(c.Param("user_id"), 10, 64))
	if err != nil {
		apiError := &utils.ApplicationError{
			Message: "User ID must me a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		c.JSON(apiError.Status, apiError)
		return
	}
	user, error := services.UsersService.GetUser(userID)
	if error != nil {
		c.JSON(error.Status, error)
		return
	}
	c.JSON(http.StatusOK, user)
}
