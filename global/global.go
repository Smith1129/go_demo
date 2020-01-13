package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis"
)

var MysqlConfig = "zzz:zzz96371@tcp(www.inlighting.org:3306)/zzz?charset=utf8&parseTime=true&loc=Local"

var GormConfig *gorm.DB

func GetRedisClient() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
	})
	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisdb
}

