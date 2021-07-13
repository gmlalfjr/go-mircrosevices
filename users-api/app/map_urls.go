package app

import (
	"github.com/go-microservices/users-api/controller"
)

func mapUrls() {
	router.POST("/user", controller.RegisterUser)
	router.POST("/user/login", controller.LoginUser)
}
