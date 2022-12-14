package app

import (
	"github.com/gin-gonic/gin"
	"github.com/menxqk/rest-microservices-in-go/common/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8080")
}
