package app

import (
	"golang-microservices/mvc/controllers"
	"net/http"
)

// StartApp initialize the routes and start the server
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
