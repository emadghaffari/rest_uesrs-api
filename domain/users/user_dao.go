package users

// data access object

import (
	"fmt"

	"github.com/emadghaffari/bookstore_uesrs-api/datasources/mysql/userdb"
	"github.com/emadghaffari/bookstore_uesrs-api/utils/date"
	"github.com/emadghaffari/bookstore_uesrs-api/utils/errors"
	"github.com/emadghaffari/bookstore_uesrs-api/utils/mysql"
)

const (
	indexUniqueEmail       = "email_UNIQUE"
	insertQyery            = "INSERT INTO users(first_name, last_name, email, password, created_at) VALUES(?, ?, ?, ?, ?);"
	selectQuery            = "SELECT id, first_name, last_name, email, created_at FROM users WHERE id=?"
	updateQuery            = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	deleteQuery            = "DELETE FROM users WHERE id=?"
	findUsersByStatusQuery = "SELECT id, status , first_name, last_name, email, created_at FROM users WHERE status=?;"
)

// Get a user if exists in DB
func (user *User) Get() *errors.ResError {
	stms, err := userdb.Client.Prepare(selectQuery)
	if err != nil {
		return mysql.ParseError(err)
	}
	defer stms.Close()
	result := stms.QueryRow(user.ID)

	if err := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
	); err != nil {
		return mysql.ParseError(err)
	}

	return nil
}

// Save a user if not exists in DB
func (user *User) Save() *errors.ResError {
	stms, err := userdb.Client.Prepare(insertQyery)
	if err != nil {
		return errors.HandlerInternalServerError(err.Error())
	}
	defer stms.Close()

	result, err := stms.Exec(user.FirstName, user.LastName, user.Email, user.Password, date.GetNowString())
	if err != nil {
		return mysql.ParseError(err)
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		return mysql.ParseError(err)
	}
	return nil
}

// Update func try to access data and update user
func (user *User) Update() *errors.ResError {
	stms, err := userdb.Client.Prepare(updateQuery)
	if err != nil {
		return errors.HandlerInternalServerError(err.Error())
	}
	defer stms.Close()
	_, err = stms.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql.ParseError(err)
	}
	return nil
}

// Delete meth
func (user *User) Delete() *errors.ResError {
	stms, err := userdb.Client.Prepare(deleteQuery)
	if err != nil {
		return mysql.ParseError(err)
	}
	defer stms.Close()

	if _, err := stms.Exec(user.ID); err != nil {
		return mysql.ParseError(err)
	}

	return nil
}

// FindByStatus func
func (user *User) FindByStatus(status string) ([]User, *errors.ResError) {
	stms, err := userdb.Client.Prepare(findUsersByStatusQuery)
	if err != nil {
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Prepare %s", err.Error()))
	}
	defer stms.Close()

	row, err := stms.Query(status)
	if err != nil {
		return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Query %s", err.Error()))
	}
	defer row.Close()
	result := make([]User, 0)
	for row.Next() {
		var ur User
		if err := row.Scan(&ur.ID, &ur.Status, &ur.FirstName, &ur.LastName, &ur.Email, &ur.CreatedAt); err != nil {
			return nil, errors.HandlerInternalServerError(fmt.Sprintf("Error in Scan %s", err.Error()))
		}
		result = append(result, ur)
	}
	return result, nil
}
