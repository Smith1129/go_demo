package global

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var MysqlConfig = "zzz:zzz96371@tcp(www.inlighting.org:3306)/zzz?charset=utf8&parseTime=true&loc=Local"

var GormConfig *gorm.DB

func GetRedisClient() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Sprintf("%v","--------redis启动失败--------",err)
	}
	return redisdb
}

