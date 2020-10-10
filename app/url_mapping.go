package app

import (
	"github.com/emadghaffari/bookstore_uesrs-api/controllers/ping"
	"github.com/emadghaffari/bookstore_uesrs-api/controllers/users"
)

func mapURL() {
	router.GET("/ping", ping.Ping)

	// users CROD
	router.POST("/user", users.CreateUser)
	router.GET("/user/:user_id", users.GetUser)
	// PUT is for full update and PATCH for parshal update(just fileds exists in body)
	router.PUT("/user/:user_id", users.UpdateUser)
}
