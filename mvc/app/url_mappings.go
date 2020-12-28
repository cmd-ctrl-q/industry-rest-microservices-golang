package app

import (
	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/controllers"
)

func mapUrls() {

	router.GET("/users/:user_id", controllers.GetUser)

}
