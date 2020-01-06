package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SophieDeBenedetto/golang-microservices/mvc/services"
	"github.com/SophieDeBenedetto/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))
	if err != nil {
		apiError := &utils.ApplicationError{
			Message: "User ID must me a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.Status)
		resp.Write(jsonValue)
		return
	}
	user, error := services.GetUser(userId)
	if error != nil {
		jsonValue, _ := json.Marshal(error)
		resp.WriteHeader(error.Status)
		resp.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(*user)
	resp.Write(jsonValue)
}
