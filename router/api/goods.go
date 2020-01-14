package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"go_demo/models"
	"go_demo/utils"
	"net/http"
	"strconv"
)

func GetGoodList(c *gin.Context){
	//var good []models.Good
	page,err := strconv.Atoi(c.DefaultQuery("Page","1"))
	if err != nil{
		fmt.Println(err)
	}
	//pageSize := 10
	result := utils.Paginator(5,page,"good")
	//fmt.Println(page,pageSize,page*pageSize)
	////分页 limit ->返回多少个数据 offset从第几个数据开始
	//global.GormConfig.Limit(10).Offset(5).Find(&good)
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":result})
}
