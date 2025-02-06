package handler

import (
	"net/http"
	"strconv"

	"github.com/MyNameIsWhaaat/todo-app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept json
// @Produce json
// @Param input body todo.TodoList true "list info"
// @Success 200 {integer} integer
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
    userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    var input todo.TodoList
    if err := c.BindJSON(&input); err != nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        logrus.Errorf("failed to bind JSON: %s", err.Error())
        return
    }

    id, err := h.services.TodoList.Create(userId, input)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        logrus.Errorf("failed to create todo list: %s", err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{
        "id": id,
    })
}

type getAllListsResponse struct{
    Data []todo.TodoList `json:"data"`
}

// @Summary Get all todo lists
// @Security ApiKeyAuth
// @Tags lists
// @Description Retrieves all todo lists for the authenticated user
// @ID get-lists
// @Accept json
// @Produce json
// @Success 200 {object} getAllListsResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    lists, err := h.services.TodoList.GetAll(userId)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        logrus.Errorf("failed to create todo list: %s", err.Error())
        return
    }

    c.JSON(http.StatusOK, getAllListsResponse{
        Data: lists,
    })
}

// @Summary Get todo list by ID
// @Security ApiKeyAuth
// @Tags lists
// @Description Retrieves a single todo list by its ID
// @ID get-list-by-id
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {object} todo.TodoList
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context){
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

    list, err := h.services.TodoList.GetById(userId, id)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        logrus.Errorf("failed to create todo list: %s", err.Error())
        return
    }

    c.JSON(http.StatusOK, list)
}

// @Summary Update todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description Updates a todo list by its ID
// @ID update-list
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Param input body todo.UpdateListInput true "Update params"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context){
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

    var input todo.UpdateListInput
    if err:= c.BindJSON(&input); err != nil{
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.services.TodoList.Update(userId, id, input)
    err != nil{
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, statusResponse{"Ok"})
}

// @Summary Delete todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description Deletes a todo list by its ID
// @ID delete-list
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context){
    userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "invalid id param")
    }

    err = h.services.TodoList.Delete(userId, id)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        logrus.Errorf("failed to create todo list: %s", err.Error())
        return
    }

    c.JSON(http.StatusOK, statusResponse{
        Status: "Ok",
    })
}