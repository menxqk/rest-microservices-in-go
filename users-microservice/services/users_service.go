package services

import (
	"fmt"

	"github.com/menxqk/rest-microservices-in-go/common/crypto"
	"github.com/menxqk/rest-microservices-in-go/common/date"
	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/menxqk/rest-microservices-in-go/users-microservice/domain/users"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	SearchUsers(string) (users.Users, *errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestError)
}

type usersService struct{}

func (us *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (us *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date.GetNowAsString()
	user.Status = users.STATUS_ACTIVE
	user.Password = crypto.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := us.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
		if user.Password != "" {
			current.Password = crypto.GetMd5(user.Password)
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
		if user.Password != "" {
			current.Password = crypto.GetMd5(user.Password)
		} else {
			current.Password = ""
		}
	}

	fmt.Printf("current: %+v\n", current)

	if err := current.Validate(); err != nil {
		return nil, err
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (us *usersService) DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (us *usersService) SearchUsers(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (us *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
