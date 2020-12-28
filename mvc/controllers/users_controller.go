package controllers

import (
	"net/http"
	"strconv"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/services"
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
)

// GetUser ...
func GetUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		// utils.Respond(c, apiErr.StatusCode, apiErr)
		// c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	// user id invalid, or internal service error (db error), etc.
	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		// utils.Respond(c, apiErr.StatusCode, apiErr)
		// c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	// user exists,
	utils.Respond(c, http.StatusOK, user)
	// c.JSON(http.StatusOK, user)
}
