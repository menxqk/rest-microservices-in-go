package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/menxqk/rest-microservices-in-go/bookstore_oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
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
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}
