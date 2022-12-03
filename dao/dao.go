package dao

import (
	sql "database/sql"
	"log"
	"message-board/model"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/messageboard?charset=uft8mb4&loc=Local&parseTime=true")
	if err != nil {
		log.Fatalf("connect mysql error:%v", err)
		return
	}
	DB = db
}
func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user(name,password)values(?,?)", u.UserName, u.Password)
	return
}
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from user where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Id, &u.UserName, &u.Password)
	return
}

func SearchMessage(username string) (u model.Message, err error) {
	row := DB.QueryRow("select mid,send_name,rec_name,details from message where rec_name=?", username)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.MID, &u.SendName, &u.Receiver, &u.Details)
	return
}

func ModifyPassword(name string, newPassword string) (err error) {
	_, err = DB.Exec("update user set password = ? where username = ?", newPassword, name)
	return
}
func LeaveMessage(sender string, receiver string, details string) (err error) {
	_, err = DB.Exec("insert into message(send_name,rec_name,details)values (?,?)", sender, receiver, details)
	return
}
func ModifyMessage(u model.Message) (err error) {
	_, err = DB.Exec("update message set (rec_name,details)=(?,?) where send_name=?", u.Receiver, u.Details, u.SendName)
	return
}
func DeleteMessage(sender string, receiver string) (err error) {
	_, err = DB.Exec("delete from message where (send_name,rec_name)=(?,?)", sender, receiver)
	return
}
func SearchComment(recMID string) (u model.Comment, err error) {
	row := DB.QueryRow("select mid,send_name,rec_name,details from comment where mid=?", recMID)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.MID, &u.RecMID, &u.Sender, &u.Details)
	return
}
func LeaveComment(recMID string, sender string, details string) (err error) {
	_, err = DB.Exec("insert into comment(rec_mid,sender,details)values (?,?,?)", recMID, sender, details)
	return
}
func ModifyComment(u model.Comment) (err error) {
	_, err = DB.Exec("update comment set (rec_mid,details)=(?,?) where sender=?", u.RecMID, u.Details, u.Sender)
	return
}
func DeleteComment(sender string, recMID string) (err error) {
	_, err = DB.Exec("delete from comment where (sender,rec_mid)=(?,?)", sender, recMID)
	return
}
