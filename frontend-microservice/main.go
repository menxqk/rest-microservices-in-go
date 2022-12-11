package main

import (
	"github.com/menxqk/dotenv"
	"github.com/menxqk/rest-microservices-in-go/frontend-microservice/app"
)

func main() {
	err := dotenv.LoadFile(".env")
	if err != nil {
		panic(err)
	}

	app.StartApplication()
}
