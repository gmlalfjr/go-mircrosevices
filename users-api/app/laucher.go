package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func RunServer() {
	mapUrls()
	router.Run(":9000")
}
