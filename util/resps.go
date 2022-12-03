package util

import "github.com/gin-gonic/gin"

type respTemplate struct {
	Status int    `json:"Status"`
	Info   string `josn:"Info"`
}

var OK = respTemplate{
	Status: 200,
	Info:   "success",
}

var ParamError = respTemplate{
	Status: 300,
	Info:   "params error",
}
var InternalError = respTemplate{
	Status: 500,
	Info:   "internal error",
}

func RespOK(c *gin.Context) {
	c.JSON(200, OK)
}

func RespParamErr(c *gin.Context) {
	c.JSON(300, ParamError)
}

func RespInternalError(c *gin.Context) {
	c.JSON(500, InternalError)
}
func NormErr(c *gin.Context, status int, info string) {
	c.JSON(300, gin.H{
		"status": 300,
		"info":   info,
	})
}
