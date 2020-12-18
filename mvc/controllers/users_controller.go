package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/services"
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
)

/*
the responsibility of a controller is to make sure we have every parameter we need

*/

// GetUser  will validate all the parameters we need before calling service.GetUser()
func GetUser(w http.ResponseWriter, r *http.Request) {

	userIDParam := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	// if there is an error, then user_id is not a number, so make a new error and return its json value back to the client
	if err != nil {
		// new
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		// marshal the error into json
		jsonValue, _ := json.Marshal(appErr)

		// return json value of error back to client
		w.WriteHeader(appErr.StatusCode)
		// w.WriteHeader(http.StatusBadRequest) // 400 bad request
		w.Write(jsonValue) // write the json to client
		return
	}

	// get user from services.GetUser
	user, apiErr := services.GetUser(userID)
	// if there is an error, then
	if apiErr != nil {
		// marshal the new error into json
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user) // format user before sending back to client
	w.Write(jsonValue)                 // send user in json format to client

}
