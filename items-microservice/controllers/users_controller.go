package controllers

import "net/http"

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
}

type usersController struct{}

func (u *usersController) Create(w http.ResponseWriter, r *http.Request) {

}
