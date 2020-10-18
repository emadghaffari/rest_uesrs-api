package users

import (
	"net/http"
	"strconv"

	"github.com/emadghaffari/rest_uesrs-api/domain/users"
	"github.com/emadghaffari/rest_uesrs-api/services"
	cryptoutils "github.com/emadghaffari/rest_uesrs-api/utils/cryptoUtils"
	"github.com/emadghaffari/rest_uesrs-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(ID string) (int64, *errors.ResError) {
	userID, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return 0, errors.HandlerBagRequest("Invalid, ID should be a Number")
	}
	return userID, nil
}

// Get func for get a users
func Get(c *gin.Context) {
	// get user_id
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	// get a new User Domain(Core of project) and return response
	result, resErr := services.GetUser(userID)
	if resErr != nil {
		c.JSON(resErr.Status, resErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall((c.GetHeader("X-Public") == "true")))
}

// Create func for store new user
func Create(c *gin.Context) {
	// create a variable from user struct
	user := users.User{}

	// Bind the request.Body to user
	if err := c.ShouldBindJSON(&user); err != nil {
		resErr := errors.HandlerBagRequest("Invalid JSON Body.")
		c.JSON(resErr.Status, resErr)
		return
	}

	user.Password = cryptoutils.GetMD5(user.Password)
	// create a new User Domain(Core of project) and return response
	result, resErr := services.CreateUser(user)
	if resErr != nil {
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall((c.GetHeader("X-Public") == "true")))
}

// Update func
func Update(c *gin.Context) {
	// create a variable from user struct
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	// create a variable from user struct
	user := users.User{}
	user.ID = userID

	// Bind the request.Body to user
	if err := c.ShouldBindJSON(&user); err != nil {
		resErr := errors.HandlerBagRequest("Invalid JSON Body.")
		c.JSON(resErr.Status, resErr)
		return
	}

	// tru to update a single user
	result, resErr := services.UpdateUser(user)
	if resErr != nil {
		c.JSON(resErr.Status, resErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall((c.GetHeader("X-Public") == "true")))
}

// Delete func
func Delete(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Status func
func Status(c *gin.Context) {
	result, err := services.GetByStatus(c.Param("status"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}
