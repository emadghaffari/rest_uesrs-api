package users

import (
	"net/http"
	"strconv"

	"github.com/emadghaffari/bookstore_uesrs-api/domain/users"
	"github.com/emadghaffari/bookstore_uesrs-api/services"
	"github.com/emadghaffari/bookstore_uesrs-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// GetUser func for get a users
func GetUser(c *gin.Context) {
	// create a variable from user struct
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		resErr := errors.HandlerBagRequest("Invalid, ID should be a Number")
		c.JSON(resErr.Status, resErr)
		return
	}

	// get a new User Domain(Core of project) and return response
	result, resErr := services.GetUser(userID)
	if resErr != nil {
		c.JSON(resErr.Status, resErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// CreateUser func for store new user
func CreateUser(c *gin.Context) {
	// create a variable from user struct
	user := users.User{}

	// Bind the request.Body to user
	if err := c.ShouldBindJSON(&user); err != nil {
		resErr := errors.HandlerBagRequest("Invalid JSON Body.")
		c.JSON(resErr.Status, resErr)
		return
	}

	// create a new User Domain(Core of project) and return response
	result, resErr := services.CreateUser(user)
	if resErr != nil {
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// FindUser func for find the user we wants
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement ME!")
}
