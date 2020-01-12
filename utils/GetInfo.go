package utils

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go_demo/models"
	"reflect"
)

func ModelToMap(obj interface{}) map[string]interface{}{
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
//获取用户信息 转userstruct
func GetUser(c *gin.Context)(models.User,error){
	//json转struct
	var user models.User
	_userInfo,ok := c.Get("userInfo")
	if !ok{
		return user,errors.New("获取失败")
	}
	err := json.Unmarshal(_userInfo.([]byte),&user)
	if err != nil{
		return user,err
	}
	return user,nil

	//json转map
	//var mapResult map[string]interface{}
	//	//err := json.Unmarshal(userInfo.([] byte), &mapResult)
	//	//if err != nil {
	//	//	fmt.Println("JsonToMapDemo err: ", err)
	//	//}
}
