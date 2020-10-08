package services

import (
	"github.com/emadghaffari/bookstore_uesrs-api/domain/users"
	"github.com/emadghaffari/bookstore_uesrs-api/utils/errors"
)

// CreateUser func in services -> validate then create a new user
func CreateUser(user users.User) (*users.User, *errors.ResError) {
	if err := user.Validate(); err != nil {
		return nil, err

	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil

}

// GetUser func in services -> validate then create a new user
func GetUser(userID int64) (*users.User, *errors.ResError) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil

}
