package utils

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"go_demo/global"
	"log"
)
var redisClient *redis.Client = global.GetRedisClient()
func SetHashRedis(key string,field string,m map[string]interface{}) error{
	data,err := json.Marshal(m)
	if err != nil{
		log.Print(err)
		return err
	}
	_,err1 := redisClient.HSet(key,field,data).Result()
	if err1 != nil{
		log.Print(err)
		return err1
	}
	return nil
}

func GetHashRedis(key string,field string) (map[string]interface{},error){
	mapResult := make(map[string] interface{})
	data,err := redisClient.HGet(key,field).Result()
	if err != nil {
		return mapResult,err
	}
	err1 := json.Unmarshal([] byte(data),&mapResult)
	return mapResult,err1
}
