package handler

import (
	"net/http"
	"strconv"

	"github.com/MyNameIsWhaaat/todo-app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Create todo list item
// @Security ApiKeyAuth
// @Tags items
// @Description Creates a new todo list item
// @ID create-item
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Param input body todo.TodoItem true "Item info"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    listId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
    }

	var input todo.TodoItem
	if err := c.BindJSON(&input); err!= nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get all todo list items by ID
// @Security ApiKeyAuth
// @Tags items
// @Description Retrieves all todo list items by its ID
// @ID get-all-items
// @Accept json
// @Produce json
// @Param list_id path int true "List ID"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{list_id}/items [get]
func (h *Handler) getAllItems(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    listId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
    }

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get todo list item by ID
// @Security ApiKeyAuth
// @Tags items
// @Description Retrieves a single todo list item by its ID
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    itemId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
    }

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update todo list item
// @Security ApiKeyAuth
// @Tags items
// @Description Updates a todo list item by its ID
// @ID update-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param input body todo.UpdateItemInput true "Update params"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid id param")
        return
    }

    var input todo.UpdateItemInput
    if err:= c.BindJSON(&input); err != nil{
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.services.TodoItem.Update(userId, id, input)
    err != nil{
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, statusResponse{"Ok"})
}

// @Summary Delete todo list item
// @Security ApiKeyAuth
// @Tags items
// @Description Deletes a todo list item by its ID
// @ID delete-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    itemId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
    }

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}