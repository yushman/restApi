package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIdKey           = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdKey, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userIdKey)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}

func getListId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid list id param")
		return 0, err
	}
	return id, err
}

func getItemtId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Params.ByName("item_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid item id param")
		return 0, err
	}
	return id, err
}
