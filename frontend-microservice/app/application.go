package app

import (
	"net/http"
)

var (
	// router = mux.NewRouter()
	router = http.NewServeMux()
)

func StartApplication() {
	mapUrls()

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
