package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication func
func StartApplication() {
	mapURL()
	router.Run(":8080")
}
