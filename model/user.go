package model

type User struct {
	Id       int64  `json:"Id"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}
