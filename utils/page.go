package utils

import (
	"fmt"
	"go_demo/global"
	"reflect"
)

func Paginator(pageSize int,currentPage int,model string,target interface{})  map[string] interface{}{
	fmt.Println(reflect.TypeOf(target),"-----------------")
	//tablename,err := json.Marshal(model)
	//fmt.Println(tablename,err)
	//获取总数据
	var count int
	mapresult := make(map[string]interface{})
	//var data interface{}
	global.GormConfig.Table(model).Count(&count)
	//分页 limit ->返回多少个数据 offset从第几个数据开始
	global.GormConfig.Limit(pageSize).Offset(currentPage*pageSize).Find(reflect.ValueOf(&target))
	mapresult["Count"] = count
	mapresult["Page"] = currentPage
	mapresult["list"] = &target
	return mapresult
}
