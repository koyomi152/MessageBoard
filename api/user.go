package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) == 0 || len(username) > 20 {
		util.RespParamErr(c)
		return
	}
	if len(password) < 6 {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(username)
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalError(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 300, "用户已存在")
		return
	}
	err = service.CreatUser(model.User{
		UserName: username,
		Password: password,
	})
	if err != nil {
		util.RespParamErr(c)
		return
	} else {
		util.RespInternalError(c)
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) == 0 || len(username) > 20 {
		util.RespParamErr(c)
		return
	}
	if len(password) < 6 {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
			return
		} else {
			log.Printf("search user error:%v", err)
		}
		return
	}
	if u.Password != password {
		util.NormErr(c, 20001, "密码错误")
		return
	}
	c.SetCookie("name", username, 0, "", "/", false, false)
}

func Password(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	newPassWord := c.PostForm("newPassword")
	u, err := service.SearchUserByUserName(username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
			return
		} else {
			log.Printf("search user error:%v", err)
		}
		return
	}
	if u.Password != password {
		util.NormErr(c, 20001, "密码错误")
		return
	}
	err = service.ModifyPassword(username, newPassWord)
	if err != nil {
		fmt.Println(err)
		return
	}
	util.RespOK(c)
}
