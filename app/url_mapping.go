package app

import (
	"github.com/emadghaffari/bookstore_uesrs-api/controllers/ping"
	"github.com/emadghaffari/bookstore_uesrs-api/controllers/users"
)

func mapURL() {
	router.GET("/ping", ping.Ping)

	// store a new user
	router.POST("/user", users.Create)
	// GET a single user with user_id
	router.GET("/user/:user_id", users.Get)
	// PUT is for full update and PATCH for parshal update(just fileds exists in body)
	router.PUT("/user/:user_id", users.Update)
	// DELETE a user with user_id
	router.DELETE("/user/:user_id", users.Delete)

	// find users by status
	router.GET("/status/:status", users.Status)
}
