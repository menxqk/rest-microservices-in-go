package app

import (
	"net/http"

	"github.com/menxqk/rest-microservices-in-go/items-microservice/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
