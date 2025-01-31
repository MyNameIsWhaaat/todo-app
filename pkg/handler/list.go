package handler

import (
	"net/http"

	"github.com/MyNameIsWhaaat/todo-app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createList(c *gin.Context) {
    userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    var input todo.TodoList
    if err := c.BindJSON(&input); err != nil {
        newErrorResponce(c, http.StatusBadRequest, err.Error())
        logrus.Errorf("failed to bind JSON: %s", err.Error())
        return
    }

    id, err := h.services.TodoList.Create(userId, input)
    if err != nil {
        newErrorResponce(c, http.StatusInternalServerError, err.Error())
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

func (h *Handler) getAllLists(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    lists, err := h.services.TodoList.GetAll(userId)
    if err != nil {
        newErrorResponce(c, http.StatusInternalServerError, err.Error())
        logrus.Errorf("failed to create todo list: %s", err.Error())
        return
    }

    c.JSON(http.StatusOK, getAllListsResponse{
        Data: lists,
    })
}

func (h *Handler) getListById(c *gin.Context){
	
}

func (h *Handler) updateList(c *gin.Context){
	
}

func (h *Handler) deleteList(c *gin.Context){
	
}