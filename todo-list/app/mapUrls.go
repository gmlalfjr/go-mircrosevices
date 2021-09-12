package app

import (
	"github.com/gmlalfjr/go_authMiddleware/auth"
	controller "github.com/go-microservices/todo-list/controllers"
)

func route() {
	router.POST("/todo", auth.VerifyAuthorization, controller.AddTodo)
	router.PATCH("/todo/:post_id", auth.VerifyAuthorization, controller.UpdateTodo)
	router.GET("/todos", auth.VerifyAuthorization, controller.GetTodo)
	router.DELETE("/todo/:post_id", auth.VerifyAuthorization, controller.DeleteTodo)
}
