package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/domain/users"
	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/services"
	"github.com/menxqk/rest-microservices-in-go/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		idErr := errors.NewBadRequestError("invalid user id")
		c.JSON(idErr.Status, idErr)
		return
	}

	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	search, _ := c.Params.Get("search")
	c.String(http.StatusNotImplemented, fmt.Sprintf("not implemented yet: SearchUser (%s)\n", search))
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
