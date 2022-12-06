package app

import (
	"github.com/gin-gonic/gin"
	"github.com/menxqk/rest-microservices-in-go/bookstore_oauth-api/src/domain/access_token"
	"github.com/menxqk/rest-microservices-in-go/bookstore_oauth-api/src/http"
	"github.com/menxqk/rest-microservices-in-go/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// create repo
	dbRepository := db.NewRepository()

	// create service
	atService := access_token.NewService(dbRepository)

	// create handler
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
