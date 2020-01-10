package global

import "github.com/jinzhu/gorm"

var MysqlConfig = "root:123456@tcp(localhost:3306)/go_demo?charset=utf8"

var GormConfig *gorm.DB
