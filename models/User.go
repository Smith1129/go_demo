package models

import (
	"errors"
	"fmt"
	"go_demo/global"
	"strings"
)

type User struct {
	ID	int	`gorm:"primary_key" json:"id"` //
	Username	string	`json:"username" binding:"required" ` //
	Pass	string	`json:"pass" binding:"required"` //
	Address	string	`json:"address"` //
	Nickname	string	`json:"nickname"` //
	Avatar	string	`gorm:"default":'aaa' json:"avatar"` //
	Money	string	`gorm:"default":0.00 json:"money"` //
}

func (u *User) InsertUser() int{
	//数组切片
	usernameStr := strings.Replace(u.Username," ","",-1)
	passStr := strings.Replace(u.Pass," ","",-1)
	if len(usernameStr) == 0 || len(passStr)==0{
		return 2
	}
	var user []User
	global.GormConfig.Where("username = ?",u.Username).Find(&user)
	if len(user) == 0{
		err := global.GormConfig.Create(u).Error
		fmt.Print(err)
		return  1
	}else {
		return 0
	}
}

func (u *User) FindUserByUsernameAndPass() ([]User,error){
	var user []User
	global.GormConfig.Where("username = ? And pass = ?",u.Username,u.Pass).Find(&user)
	if len(user) == 0{
		return user,errors.New("账号密码错误")
	}
	return user,nil
}
func (u *User) FindUserByUsername() []User{
	var user []User
	global.GormConfig.Where("username = ?",u.Username).Find(&user)
	return user
}


