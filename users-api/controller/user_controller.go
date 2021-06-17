package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/users-api/domain"
	"github.com/users-api/services"
)

func RegisterUser(c *gin.Context) {
	user := domain.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, err := services.RegisterUser(&user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func LoginUser(c *gin.Context) {
	userLogin := domain.UserLogin{}
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	result, err := services.LoginUser(&userLogin)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
