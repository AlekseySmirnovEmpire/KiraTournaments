package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")

		return
	}

	headerPath := strings.Split(header, " ")
	if len(headerPath) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")

		return
	}

	uID, err := h.services.Authorization.ParseToken(headerPath[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	c.Set(userCtx, uID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user ID not found in context")

		return 0, errors.New("user ID not found in context")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user ID has no right type")

		return 0, errors.New("user ID has no right type")
	}

	return idInt, nil
}
