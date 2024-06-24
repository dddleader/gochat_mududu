package dao

import (
	"errors"
	"gochat_mududu/db"
	"time"
)

var dbIns = db.GetDb("go_chat")

// use tag for gorm operation
type User struct {
	Id         int       `gorm:"primary_key;column:id;autoIncrement"`
	UserName   string    `gorm:"column:user_name;size:20;not null;unique;default:''"`
	Password   string    `gorm:"column:password;size:40;not null;default:''"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP"`
	db.DbGoChat
}

//operation to db

// only one table so far
func (u *User) TableName() string {
	return "user"
}

// add current user to db
func (u *User) Add() (userId int, err error) {
	if u.UserName == "" || u.Password == "" {
		return 0, errors.New("user_name or password empty!")
	}
	oUser := u.GetUserByUserName(u.UserName)
	if oUser.Id > 0 {
		return oUser.Id, nil
	}
	u.CreateTime = time.Now()
	if err = dbIns.Table(u.TableName()).Create(&u).Error; err != nil {
		return 0, err
	}
	return u.Id, nil
}

// get userdata by username
func (u *User) GetUserByUserName(userName string) (data User) {
	//use gorm to get user data
	dbIns.Table(u.TableName()).Where("user_name=?", userName).Take(&data)
	return
}

func (u *User) GetUserNameByUserId(userId int) (userName string) {
	var data User
	dbIns.Table(u.TableName()).Where("id=?", userId).Take(&data)
	return data.UserName
}

// get user by name
func (u *User) CheckHaveUserName(userName string) (data User) {
	dbIns.Table(u.TableName()).Where("user_name=?", userName).Take(&data)
	return
}
