package v1

import (
	"Project1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)

	code := model.CreateUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
	})
}
