package users

import (
	"encoding/json"
	"strings"

	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/utils/errors"
)

const (
	STATUS_ACTIVE = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (u *User) Validate() *errors.RestError {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if u.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}

func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id          int64  `json:"id"`
		DateCreated string `json:"date_created"`
		Status      string `json:"status"`
	}{
		Id:          u.Id,
		DateCreated: u.DateCreated,
		Status:      u.Status,
	})
}
