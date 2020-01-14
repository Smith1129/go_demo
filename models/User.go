package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"go_demo/global"
	"strings"
	"time"
	//"unicode/utf8"
)

type User struct {
	ID	int	`gorm:"primary_key" json:"id"` //
	Username	string	`json:"username" binding:"required" ` //
	Pass	string	`json:"-" binding:"required"` //
	Address	string	`json:"address"` //
	Nickname	string	`json:"nickname"` //
	Avatar	string	`gorm:"default":'aaa' json:"avatar"` //
	Money	float64	`gorm:"default":0.00 json:"money"` //
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func Encrypt(pass string) string{
	md5 := md5.New()
	md5.Write([] byte(pass))
	fmt.Println(hex.EncodeToString(md5.Sum(nil)),"-------")
	return hex.EncodeToString(md5.Sum(nil))
}

func (u *User) InsertUser() int{
	//数组切片
	usernameStr := strings.Replace(u.Username," ","",-1)
	passStr := strings.Replace(u.Pass," ","",-1)
	if len(usernameStr) <= 2 || len(passStr)==0{
		return 2
	}
	var user []User
	global.GormConfig.Where("username = ?",u.Username).Find(&user)
	if len(user) == 0{
		u.Pass = Encrypt(u.Pass)
		err := global.GormConfig.Create(u).Error
		if err != nil{
			return 0
		}
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
func  FindUserByUsername(username string) User{
	var user User
	global.GormConfig.Where("username = ?",username).Find(&user)
	return user
}

func (u *User) UpdateUserNameAndAddress() string{
	err := global.GormConfig.Model(&u).Updates(map[string] interface{}{"nickname":u.Nickname,"address":u.Address}).Error
	if err != nil{
		return "修改失败"
	}
	return "修改成功"
}


