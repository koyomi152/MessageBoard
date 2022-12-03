package service

import (
	"message-board/dao"
	"message-board/model"
)

func SearchComment(recMid string) (u model.Comment, err error) {
	u, err = dao.SearchComment(recMid)
	return
}
func LeaveComment(u model.Comment) error {
	err := dao.LeaveComment(u.RecMID, u.Sender, u.Details)
	return err
}
func ModifyComment(u model.Comment) error {
	err := dao.ModifyComment(u)
	return err
}
func DeleteComment(sender string, recMID string) error {
	err := dao.DeleteComment(sender, recMID)
	return err
}
