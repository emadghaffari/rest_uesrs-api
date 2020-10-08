package app

import (
	"github.com/emadghaffari/bookstore_uesrs-api/controllers/ping"
	"github.com/emadghaffari/bookstore_uesrs-api/controllers/users"
)

func mapURL() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", users.GetUser)

	router.POST("/user", users.CreateUser)
}
