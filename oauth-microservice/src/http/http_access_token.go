package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/menxqk/rest-microservices-in-go/common/errors"
	"github.com/menxqk/rest-microservices-in-go/oauth-microservice/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

type accessTokenHandler struct {
	service access_token.Service
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := c.Param("access_token_id")

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, at)
}
