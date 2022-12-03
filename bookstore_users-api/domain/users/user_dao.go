package users

import (
	"fmt"

	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/utils/date"
	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestError {
	user := usersDB[u.Id]
	if user == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}

	u.Id = user.Id
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.DateCreated = user.DateCreated

	return nil
}

func (u *User) Save() *errors.RestError {
	user := usersDB[u.Id]
	if user != nil {
		if user.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.Id))
	}

	u.DateCreated = date.GetNowAsString()
	usersDB[u.Id] = u

	return nil
}
