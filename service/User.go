package service

import (
	"fmt"
	"go_demo/global"
	"go_demo/models"
)

func LoginCheck(u models.User)  models.User {
	err := global.GormConfig.Find(u)
	if err != nil{
		fmt.Println("worry")
	}
	return u
}
