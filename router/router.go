package router

import (
	v1 "Project1/api/v1"
	"Project1/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()

	rgPublic := r.Group("/api/v1/public")
	{
		user := rgPublic.Group("user")
		{
			user.POST("add", v1.AddUser)
			user.POST("login", v1.Login)
		}
	}
	rgPrivate := r.Group("/api/v1/private")
	rgPrivate.Use(middleware.JwtToken())
	{
		article := rgPrivate.Group("article")
		{
			article.POST("create", v1.CreateArticle)
			article.GET("get", v1.GetArticleByTitle)
		}
	}
	_ = r.Run(":4911")
}
