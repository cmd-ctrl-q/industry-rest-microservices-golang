package app

import (
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/controllers/polo"
	"github.com/cmd-ctrl-q/industry-rest-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	// any requests coming in from /repositories is going to be handled by the repositories controller
	router.POST("/repositories", repositories.CreateRepo)
}
