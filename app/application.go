package app

import (
	"github.com/emadghaffari/bookstore_uesrs-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	mapURL()
	logger.Log.Info("about to start application")
	router.Run(":8080")
}
