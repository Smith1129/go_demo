package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_demo/models"
	"net/http"
)

func Search(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//u.Login()
	 //err := c.ShouldBindJSON(&t)
	 //if err != nil{
		// c.JSON(http.StatusBadRequest, gin.H{"error": "username参数必须"})
		// return
	 //}
	//var user models.User
	u.Login()
	c.JSON(http.StatusOK, gin.H{
		"code":  "200",
		"data":  u,
	})
}

func Register( c *gin.Context){
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Code": "0","Msg":"参数不正确"})
		return
	}
	result := u.InsertUser()
	if result == 1{
		c.JSON(http.StatusOK, gin.H{"Code": "1","Data":"注册成功"})
	}else if result == 0{
		c.JSON(http.StatusOK, gin.H{"Code": "0","Msg":"注册失败,存在相同用户名"})
	}else if result == 2{
		c.JSON(http.StatusOK, gin.H{"Code": "0","Msg":"注册失败,账号密码不能为空"})
	}
}

func Test(c *gin.Context){
	var payload models.UserClaims
	payload.ID = 1
	payload.Username = "zzz"
	data,err := models.CreateToken(payload)
	fmt.Print(data,err)
	//payload.id = 1
	//payload.username = "zzz"
	//data,err := CreateToken(payload)
	//fmt.Print(data,err,"zzz")
	//CreateToken
}

