package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-microservices/todo-list/domain"
	"github.com/go-microservices/todo-list/services"
	restErr "github.com/go-microservices/todo-list/utils"
)

func AddTodo(c *gin.Context) {
	todo := &domain.Todo{}
	userId, _ := c.Get("id")

	str := fmt.Sprint(userId)

	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadGateway, restErr.NewBadRequest(fmt.Sprintf("error when hit api: %s", err)))
		return
	}

	if err := todo.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, restErr.NewBadRequest(fmt.Sprintf("error when hit api: %s", err)))
		return
	}

	result, err := services.AddTodoService(todo, str)
	if err != nil {
		c.JSON(http.StatusBadRequest, restErr.NewBadRequest(fmt.Sprintf("Error Create Data: %s", err.Message)))
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateTodo(c *gin.Context) {
	todo := &domain.Todo{}
	userId, _ := c.Get("id")
	idTodo, userErr := strconv.ParseInt(c.Param("post_id"), 10, 64)
	if userErr != nil {
		err := restErr.NewBadRequest(("invalid post id"))
		c.JSON(err.Status, err)
		return
	}
	str := fmt.Sprint(userId)
	todo.UserId = str
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadGateway, restErr.NewBadRequest(fmt.Sprintf("error when hit api: %s", err)))
		return
	}
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateTodoService(todo, idTodo, isPartial)
	if err != nil {
		c.JSON(http.StatusBadRequest, restErr.NewBadRequest(fmt.Sprintf("Error Update Data: %s", err.Message)))
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetTodo(c *gin.Context) {
	userId, _ := c.Get("id")
	str := fmt.Sprint(userId)
	todo := &domain.Todo{
		UserId: str,
	}

	result, err := services.GetTodoService(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, restErr.NewBadRequest(fmt.Sprintf("Error GET Data: %s", err.Message)))
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteTodo(c *gin.Context) {
	idTodo, userErr := strconv.ParseInt(c.Param("post_id"), 10, 64)
	if userErr != nil {
		err := restErr.NewBadRequest(("invalid post id"))
		c.JSON(err.Status, err)
		return
	}

	err := services.DeleteTodoService(idTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest, restErr.NewBadRequest(fmt.Sprintf("Error Delete Data: %s", err.Message)))
		return
	}

	c.JSON(http.StatusOK, "Sueccess")
}
