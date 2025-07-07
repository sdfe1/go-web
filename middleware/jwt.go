package middleware

import (
	"Project1/message"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(viper.GetString("mode.JwtKey"))
var code int

// MyClaims 自定义声明类型 并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username           string `json:"username"`
	jwt.StandardClaims        //内嵌标准的声明
}

// token生成过程
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	//setclaims
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),                 //过期时间
			Issuer:    viper.GetString("mode.JwtIssuer"), //签发人
		},
	}

	//用指定的签名方法创建签名对象
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	//生成签名字符串
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return " ", message.ERROR
	}
	return token, message.SUCCESS

}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, message.SUCCESS
	} else {
		return nil, message.ERROR
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenHeader := c.Request.Header.Get("Authorization")
		code = message.SUCCESS

		//判断token是否存在
		if tokenHeader == "" {
			code = message.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message.GetMsg(code),
			})

			c.Abort()
			return
		}

		//判断token是否格式正确
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer:" {
			code = message.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message.GetMsg(code),
			})
			c.Abort()
			return
		}

		//判断token是否输入正确
		key, tCode := CheckToken(checkToken[1])
		if tCode == message.ERROR {
			code = message.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message.GetMsg(code),
			})
			c.Abort()
			return
		}

		//验证token是否过期
		if time.Now().Unix() > key.ExpiresAt {
			code = message.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
