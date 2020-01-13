package tasks

import (
	"fmt"
	"go_demo/global"
	"go_demo/models"
	"time"
)
//定时任务 每隔1几分钟执行一次
func GetUserInfo(f func()){
	var ch chan int
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for range ticker.C{
			f()
		}
		ch <- 1
	}()
	<-ch
}

func TasksTest() func(){
	return func(){
		key := "userList"
		redisClient := global.GetRedisClient()
		if redisClient == nil {
			fmt.Errorf("StringDemo redisClient is nil")
			return
		}
		var user []models.User
		 global.GormConfig.Find(&user)
		fmt.Println(key,redisClient)
		//if val == nil{
		//	fmt.Println("worry")
		//}
		//fmt.Println(val.Val())
	}
}
