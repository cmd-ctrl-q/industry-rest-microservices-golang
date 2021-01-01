package repositories

import (
	"net/http"

	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/domain/repositories"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/services"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	// T1: check if json in request is valid
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr) // return bad request as a json
		return
	}
	// T2: result, err := respone of the repositories_service
	// use json to populate json request
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err) // return error status and the error
		return
	}

	// T3:
	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	// T1: check if json in request is valid
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr) // return bad request as a json
		return
	}
	// T2: result, err := respone of the repositories_service
	// use json to populate json request
	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err) // return error status and the error
		return
	}

	// T3:
	// c.JSON(http.StatusCreated, result)
	c.JSON(result.StatusCode, result)
}
