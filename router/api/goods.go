package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_demo/models"

	//"go_demo/models"
	"go_demo/utils"
	"net/http"
	"strconv"
)

func GetGoodList(c *gin.Context){
	var good []models.Good
	page,err := strconv.Atoi(c.DefaultQuery("Page","1"))
	if err != nil{
		fmt.Println(err)
		return
	}
	result := utils.Paginator(10,page,"good",good)
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":result})
}
