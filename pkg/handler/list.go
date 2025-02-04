package handler

import (
	"net/http"
	"strconv"

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
    userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponce(c, http.StatusBadRequest, "invalid id param")
        return
    }

    list, err := h.services.TodoList.GetById(userId, id)
    if err != nil {
        newErrorResponce(c, http.StatusInternalServerError, err.Error())
        logrus.Errorf("failed to create todo list: %s", err.Error())
        return
    }

    c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context){
	userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponce(c, http.StatusBadRequest, "invalid id param")
        return
    }

    var input todo.UpdateListInput
    if err:= c.BindJSON(&input); err != nil{
        newErrorResponce(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.services.Update(userId, id, input)
    err != nil{
        newErrorResponce(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, statusResponse{"Ok"})
}

func (h *Handler) deleteList(c *gin.Context){
    userId, err := getUserId(c)
    if err != nil {
        logrus.Errorf("failed to get user id: %s", err.Error())
        return
    }

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponce(c, http.StatusBadRequest, "invalid id param")
    }

    err = h.services.TodoList.Delete(userId, id)
    if err != nil {
        newErrorResponce(c, http.StatusInternalServerError, err.Error())
        logrus.Errorf("failed to create todo list: %s", err.Error())
        return
    }

    c.JSON(http.StatusOK, statusResponse{
        Status: "Ok",
    })
}