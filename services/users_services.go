package services

import (
	"github.com/emadghaffari/res_errors/errors"
	"github.com/emadghaffari/rest_uesrs-api/domain/users"
)

// CreateUser func in services -> validate then create a new user
func CreateUser(user users.User) (*users.User, errors.ResError) {
	if err := user.Validate(); err != nil {
		return nil, err

	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil

}

// GetUser func in services -> validate then create a new user
func GetUser(userID int64) (*users.User, errors.ResError) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil

}

// UpdateUser single user and connect to domain
func UpdateUser(user users.User) (*users.User, errors.ResError) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if user.FirstName != "" {
		current.FirstName = user.FirstName
	}
	if user.LastName != "" {
		current.LastName = user.LastName
	}
	if user.Email != "" {
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser user
func DeleteUser(userID int64) errors.ResError {
	user := &users.User{ID: userID}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

// GetByStatus func
func GetByStatus(status string) (users.Users, errors.ResError) {
	user := &users.User{Status: status}
	res, err := user.FindByStatus(user.Status)
	if err != nil {
		return nil, err
	}
	return res, nil
}
