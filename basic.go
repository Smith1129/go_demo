package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_demo/router"
	"go_demo/global"
)


func main() {
	var err error
	global.GormConfig,err = gorm.Open("mysql",global.MysqlConfig)
	if err !=nil {
		fmt.Printf("failed to create gorm mysql engine, %v", err)
	}
	global.GormConfig.DB().SetMaxIdleConns(1000)
	global.GormConfig.DB().SetMaxOpenConns(5000)
	global.GormConfig.LogMode(true)
	global.GormConfig.SingularTable(true)
	r := router.RouterConfig()
	r.Run()

}


