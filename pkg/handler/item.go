package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := getListId(c)
	if err != nil {
		return
	}
	var input restApi.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.TodoItem.CreateItem(userId, listId, input)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type allItemsResponse struct {
	Data []restApi.TodoItem
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := getListId(c)
	if err != nil {
		return
	}

	items, err := h.services.TodoItem.GetAllItems(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, allItemsResponse{Data: items})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := getItemtId(c)
	if err != nil {
		return
	}
	item, err := h.services.TodoItem.GetItemById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := getItemtId(c)
	if err != nil {
		return
	}
	if err := h.services.TodoItem.DeleteItem(userId, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := getItemtId(c)
	if err != nil {
		return
	}
	var update restApi.UpdateItemInput
	if err := c.BindJSON(&update); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.UpdateItem(userId, id, update); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
