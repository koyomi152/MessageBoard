package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"message-board/service"
	"message-board/util"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("name"); err == nil {

			u, err := service.SearchUserByUserName(cookie)
			if err != nil {
				log.Printf("search user error:%v", err)
			}
			if cookie == u.UserName {
				c.Next()
				return
			}
		}
		util.NormErr(c, 300, "用户未登录")
		c.Abort()
		return
	}
}
