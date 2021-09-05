package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jgmc3012/bookstore_oauth-api/src/domain/access_token"
	"github.com/jgmc3012/bookstore_oauth-api/src/services"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service services.AccessTokenService
}

func NewHandler(service services.AccessTokenService) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request access_token.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}
	at, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}
