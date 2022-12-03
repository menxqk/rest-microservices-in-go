package app

import (
	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/controllers/ping"
	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search/:search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
