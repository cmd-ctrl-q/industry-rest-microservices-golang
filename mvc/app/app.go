package app

import (
	"net/http"

	"github.com/cmd-ctrl-q/industry-rest-microservices/mvc/controllers"
)

func StartApp() {

	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
