package mysql

import (
	"strings"

	"github.com/emadghaffari/res_errors/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

// ParseError func for mysql errors
func ParseError(err error) errors.ResError {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.HandlerNotFoundError("no record mached")
		}
		return errors.HandlerInternalServerError("error parsing data to database", err)
	}

	switch sqlError.Number {
	case 1062:
		return errors.HandlerInternalServerError("invalid dublicated data", err)
	}
	return errors.HandlerInternalServerError("error in proccessing request", err)
}
