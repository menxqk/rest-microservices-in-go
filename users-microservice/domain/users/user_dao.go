package users

import (
	"fmt"
	"strings"

	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/menxqk/rest-microservices-in-go/common/logger"
	"github.com/menxqk/rest-microservices-in-go/users-microservice/datasources/mysql/users_db"
	"github.com/menxqk/rest-microservices-in-go/users-microservice/utils/mysql_utils"
)

const (
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status, password from users WHERE id=?;"
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

func (u *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("error when trying to get user", errors.NewError("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)

	err = result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status, &u.Password)
	if err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("error when trying to get user", errors.NewError("database error"))
	}

	return nil
}

func (u *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("error when trying to save user", errors.NewError("database error"))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password)
	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("error when trying to save user", errors.NewError("database error"))
	}

	u.Id, err = insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("error when trying to save user", errors.NewError("database error"))
	}

	return nil
}

func (u *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("error when trying to update user", errors.NewError("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("error when trying to update user", errors.NewError("database error"))
	}

	return nil
}

func (u *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("error when trying to delete user", errors.NewError("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id)
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("error when trying to delete user", errors.NewError("database error"))
	}

	return nil
}

func (u *User) FindByStatus(status string) (Users, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by statys statement", err)
		return nil, errors.NewInternalServerError("error when trying to find users", errors.NewError("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users", err)
		return nil, errors.NewInternalServerError("error when trying to find users", errors.NewError("database error"))
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password)
		if err != nil {
			logger.Error("error when trying to scan user row into user struct", err)
			return nil, errors.NewInternalServerError("error when trying to find users", errors.NewError("database error"))
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status '%s'", status))
	}

	return results, nil
}

func (u *User) FindByEmailAndPassword() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("error when trying to find user", errors.NewError("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Email, u.Password, STATUS_ACTIVE)

	err = result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status)
	if err != nil {
		if strings.Contains(err.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", err)
		return errors.NewInternalServerError("error when trying to find user", errors.NewError("database error"))
	}

	return nil
}
