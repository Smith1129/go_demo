package utils

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go_demo/models"
	//"reflect"
)
//map转结构体
//func MapToStruct(model *interface{}) interface{}{
//	mapstruct
//}
//获取用户信息 转userstruct
func GetUser(c *gin.Context)(map[string]interface{},error){
	//json转struct
	var user models.User
	mapresult := make(map[string]interface{})
	_userInfo,ok := c.Get("userInfo")
	if !ok{
		return mapresult,errors.New("获取失败")
	}
	err := json.Unmarshal(_userInfo.([]byte),&user)
	if err != nil{
		return mapresult,err
	}
	//mapresult := make(map[string]interface{})
	mapresult["userInfo"] = user
	return mapresult,nil
}
