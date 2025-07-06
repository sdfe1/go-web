package router

import (
	v1 "Project1/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()
	rgPublic := r.Group("/api/v1/public")
	{
		user := rgPublic.Group("user")
		{
			user.POST("add", v1.AddUser)
		}
	}
	_ = r.Run(":4911")

}
