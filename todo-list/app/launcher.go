package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Launcher() {
	route()
	router.Run()
}