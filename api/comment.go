package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func GetComment(c *gin.Context) {
	recMID := c.PostForm("recMID")
	u, err := service.SearchComment(recMID)
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalError(c)
		return
	}
	if u.Details == "" {
		util.NormErr(c, 300, "无留言")
	}
	c.JSON(300, u)
}
func LeaveComment(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		util.NormErr(c, 300, "用户未登录")
	}
	recMID := c.PostForm("recMID")
	details := c.PostForm("details")
	sender := cookie
	err = service.LeaveComment(model.Comment{
		RecMID:  recMID,
		Sender:  sender,
		Details: details,
	})
	if err != nil {
		util.RespParamErr(c)
		return
	} else {
		util.RespInternalError(c)
	}
	util.RespOK(c)
	fmt.Println("已留言")
}
func ModifyComment(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		util.NormErr(c, 300, "用户未登录")
	}
	recMID := c.PostForm("recMID")
	details := c.PostForm("details")
	sender := cookie
	err = service.ModifyComment(model.Comment{
		RecMID:  recMID,
		Sender:  sender,
		Details: details,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	util.RespOK(c)

}
func DeleteComment(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		util.NormErr(c, 300, "用户未登录")
	}
	recMID := c.PostForm("recMID")
	sender := cookie
	err = service.DeleteComment(sender, recMID)
	if err != nil {
		util.RespParamErr(c)
	}
	util.RespOK(c)
}
