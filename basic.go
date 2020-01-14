package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_demo/global"
	"go_demo/router"
	"go_demo/utils"
)


func main() {
	var err error
	global.GormConfig,err = gorm.Open("mysql",global.MysqlConfig)
	if err !=nil {
		utils.Error(fmt.Sprintf("%v","-----------mysql启动失败----------",err))
	}
	global.GormConfig.DB().SetMaxIdleConns(1000)
	global.GormConfig.DB().SetMaxOpenConns(5000)
	global.GormConfig.LogMode(true)
	global.GormConfig.SingularTable(true)
	r := router.RouterConfig()
	//tasks.GetUserInfo(tasks.TasksTest())
	r.Run()
}


