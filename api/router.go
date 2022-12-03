package api

import (
	"github.com/gin-gonic/gin"
	"message-board/middleware"
)

func InitRouter() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register)
		u.POST("/login", Login)
		u.PUT("/password", middleware.AuthMiddleWare(), Password)
	}
	m := r.Group("/messsage")
	{
		m.Use(middleware.AuthMiddleWare())
		m.GET("/GetMessage", GetMessage)
		m.POST("/LeaveMessage", LeaveMessage)
		m.PUT("/ModifyMessage", ModifyMessage)
		m.DELETE("/DeleteMessage", DeleteMessage)
	}
	n := r.Group("/comment")
	{
		n.Use(middleware.AuthMiddleWare())
		n.GET("/GetComment", GetComment)
		n.POST("/LeaveComment", LeaveComment)
		n.PUT("/ModifyComment", ModifyComment)
		n.DELETE("/DeleteComment", DeleteComment)
	}
	r.Run()
}
