package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/services"
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
)

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {

	userIDParam := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)

	if err != nil {

		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(appErr)

		w.WriteHeader(appErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
