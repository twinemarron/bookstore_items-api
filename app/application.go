package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrl()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8070",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
