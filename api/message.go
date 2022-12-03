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

func GetMessage(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		util.NormErr(c, 300, "用户未登录")
	}
	u, err := service.SearchMessage(cookie)
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalError(c)
		return
	}
	if u.Details == "" {
		util.NormErr(c, 300, "无留言")
	}
	c.JSON(300, u)
}
func LeaveMessage(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		util.NormErr(c, 300, "用户未登录")
	}
	receiver := c.PostForm("receiver")
	details := c.PostForm("details")
	sender := cookie
	if len(receiver) == 0 || len(receiver) > 20 {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(receiver)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
			return
		} else {
			log.Printf("search user error:%v", err)
		}
		return
	}
	err = service.LeaveMessage(model.Message{
		SendName: sender,
		Receiver: receiver,
		Details:  details,
	})
	if err != nil {
		util.RespParamErr(c)
		return
	} else {
		util.RespInternalError(c)
	}
	util.RespOK(c)
	fmt.Println("已发送给" + u.UserName)
}
func ModifyMessage(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		util.NormErr(c, 300, "用户未登录")
	}
	receiver := c.PostForm("receiver")
	details := c.PostForm("details")
	sender := cookie
	if len(receiver) == 0 || len(receiver) > 20 {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(receiver)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
			return
		} else {
			log.Printf("search user error:%v", err)
		}
		return
	}
	err = service.ModifyMessage(model.Message{
		SendName: sender,
		Receiver: u.UserName,
		Details:  details,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	util.RespOK(c)
}
func DeleteMessage(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		util.NormErr(c, 300, "用户未登录")
	}
	receiver := c.PostForm("receiver")
	sender := cookie
	if len(receiver) == 0 || len(receiver) > 20 {
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(receiver)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NormErr(c, 300, "用户不存在")
			return
		} else {
			log.Printf("search user error:%v", err)
		}
		return
	}
	err = service.DeleteMessage(sender, u.UserName)
	if err != nil {
		util.RespParamErr(c)
	}
	util.RespOK(c)
}
