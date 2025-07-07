package v1

import (
	"Project1/message"
	"Project1/middleware"
	"Project1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 完善用户功能，注册登录，注销，
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)

	code := model.CreateUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
	})
}

// 登录功能
func Login(c *gin.Context) {
	var user model.User
	_ = c.ShouldBind(&user)

	//判断是否存在账号
	code := model.CheckLogin(&user)
	if code != message.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": message.GetMsg(code),
		})
	}

	//生成token传给客户端
	token, tokenCode := middleware.SetToken(user.Username)
	if tokenCode != message.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  tokenCode,
			"message": message.GetMsg(tokenCode),
		})
	}

	//返回成功
	c.JSON(http.StatusOK, gin.H{
		"status":  message.SUCCESS,
		"message": "登录成功",
		"token":   token,
	})
}
