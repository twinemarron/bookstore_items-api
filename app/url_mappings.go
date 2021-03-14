package app

import (
	"net/http"

	"github.com/twinemarron/bookstore_items-api/controllers"
)

func mapUrl() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
