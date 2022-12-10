package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/menxqk/rest-microservices-in-go/common/errors"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error when trying to parse error", errors.NewError("error parsing database response"))
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error when trying to parse error", errors.NewError("error processing request"))
}
