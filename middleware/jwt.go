package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		//var data interface{}
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 400
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403 //无权限的token，假
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //token失效
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
