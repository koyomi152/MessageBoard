package service

import (
	"message-board/dao"
	"message-board/model"
)

func SearchMessage(username string) (u model.Message, err error) {
	u, err = dao.SearchMessage(username)
	return
}
func LeaveMessage(u model.Message) error {
	err := dao.LeaveMessage(u.SendName, u.Receiver, u.Details)
	return err
}
func ModifyMessage(u model.Message) error {
	err := dao.ModifyMessage(u)
	return err
}
func DeleteMessage(sender string, receiver string) error {
	err := dao.DeleteMessage(sender, receiver)
	return err
}
