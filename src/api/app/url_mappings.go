package app

import (
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/controllers/polo"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	// any requests coming in from /repositories is going to be handled by the repositories controller
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
