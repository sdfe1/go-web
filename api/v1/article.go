package v1

import (
	"Project1/message"
	"Project1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)

	code := model.CreateArticle(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": message.GetMsg(code),
	})

}

func GetArticleByTitle(c *gin.Context) {
	//从查询参数中获取书名
	title := c.Query("article_title")
	//校验传入信息
	article, code := model.GetArticleByTitle(title)
	if code != message.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": message.GetMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": message.GetMsg(code),
		"data":    article,
	})

}
