package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreatUser(u model.User) error {
	err := dao.InsertUser(u)
	return err
}

func SearchUserByUserName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}
func ModifyPassword(name string, newpassword string) error {
	err := dao.ModifyPassword(name, newpassword)
	return err
}
