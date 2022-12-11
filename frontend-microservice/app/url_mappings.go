package app

import (
	"errors"
	"net/http"
	"os"

	"github.com/menxqk/rest-microservices-in-go/frontend-microservice/handlers"
	"github.com/menxqk/rest-microservices-in-go/frontend-microservice/middleware"
)

const (
	PUBLIC_HTML_DIR = "PUBLIC_HTML_DIR"
	APP_HTML_DIR    = "APP_HTML_DIR"
	APP_URL         = "APP_URL"
	AUTH_URL        = "AUTH_URL"
)

func mapUrls() {
	publicHtmlHandler := handlers.NewLocalStaticHandler(os.Getenv(PUBLIC_HTML_DIR))
	appHtmlHandler := http.StripPrefix("/app", handlers.NewLocalStaticHandler(os.Getenv(APP_HTML_DIR)))
	authHandler := handlers.NewAuthHandler()

	appUrl := os.Getenv(APP_URL)
	authUrl := os.Getenv(AUTH_URL)
	if appUrl == "" || authUrl == "" {
		panic(errors.New("invalid APP_URL or AUTH_URL"))
	}

	router.Handle("/", publicHtmlHandler)
	router.Handle(appUrl, middleware.RequireOAuth(appHtmlHandler))
	router.Handle(authUrl, authHandler)

}
