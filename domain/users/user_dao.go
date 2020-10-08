package users

// data access object

import (
	"fmt"

	"github.com/emadghaffari/bookstore_uesrs-api/utils/errors"
)

var userDB = make(map[int64]*User)

// Get a user if exists in DB
func (user *User) Get() *errors.ResError {
	result := userDB[user.ID]
	if result == nil {
		return errors.HandlerNotFoundError(fmt.Sprintf("the user %d not fount", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

// Save a user if not exists in DB
func (user *User) Save() *errors.ResError {
	if result := userDB[user.ID]; result != nil {
		if result.Email == user.Email {
			return errors.HandlerBagRequest(fmt.Sprintf("the email:%s already registerd", user.Email))
		}
		return errors.HandlerBagRequest(fmt.Sprintf("the user with this ID:%d exists in database", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
