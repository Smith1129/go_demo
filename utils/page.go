package utils

import (
	"fmt"
	"go_demo/global"
	"go_demo/models"
)

func Paginator(pageSize int,currentPage int,model string,target interface{})  map[string] interface{}{
	//获取总数据
	var count int
	mapresult := make(map[string]interface{})
	switch v := target.(type) {
	case string:
		fmt.Println("----string", v)
	case int32,int64:
		fmt.Println("-----int",v)
	case [] models.Good:
		result,_ := target.([] models.Good)
		global.GormConfig.Table(model).Count(&count)
		//分页 limit ->返回多少个数据 offset从第几个数据开始
		//order排序
		global.GormConfig.Order("like_sum desc").Limit(pageSize).Offset((currentPage-1)*pageSize).Find(&result)
		mapresult["Count"] = count
		mapresult["Page"] = currentPage
		mapresult["List"] = &result
		return mapresult
	default:
		fmt.Println("unknow")
	}
	return mapresult
}
func getData(){

}
