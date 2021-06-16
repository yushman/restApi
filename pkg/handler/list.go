package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userIdKey)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return
	}
	var input restApi.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	h.services.CreateList(int(id), input)
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}
