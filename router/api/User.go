package api

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_demo/global"
	"go_demo/models"
	"go_demo/utils"
	"net/http"
	"reflect"
	"time"
)

func Search(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  "200",
		"data":  u,
	})
}

func Register( c *gin.Context){
	var u models.User
	if err := c.BindQuery(&u); err != nil{
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

func Login(c *gin.Context){
	var u models.User
	if query_err := c.BindQuery(&u); query_err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Code": "0","Msg":"参数不正确"})
		return
	}
	u.Pass = models.Encrypt(u.Pass)
	data,find_err := u.FindUserByUsernameAndPass()
	if find_err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Code": "0","Msg":find_err.Error()})
		return
	}
	payload := models.UserClaims{
		ID:             data[0].ID,
		Username:       data[0].Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 0), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时  一小时+3600
			Issuer:    "zzz",                   //签名的发行者
		},
	}
	token,token_err := models.CreateToken(payload)
	if token_err != nil{
		c.JSON(http.StatusOK, gin.H{"Code": "0","Msg":"token创建失败"})
	}
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":token})
}

func GetUserInfo(c *gin.Context){
	userInfo,err := utils.GetUser(c)
	if err != nil{
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":userInfo.ID})
}
func Test2(c *gin.Context){
	redisClient := global.GetRedisClient()
	if redisClient == nil {
		fmt.Errorf("StringDemo redisClient is nil")
		return
	}
	var user []models.User
	global.GormConfig.Find(&user)
	mapresult := make(map[string]interface{})
	mapresult["userList"] = user
	mapresult["color"] = "red"
	data,err1 := json.Marshal(mapresult)
	fmt.Println(data,err1)
	_,err := redisClient.HSet("zzz","aa",data).Result()
	if err !=nil{
		fmt.Println(err)
	}
	ddd,err100 := redisClient.HGet("zzz","aa").Result()
	fmt.Println(ddd,err100,reflect.TypeOf(ddd))
	map1 := make(map[string] interface{})
	err11 := json.Unmarshal([] byte(ddd),&map1)
	fmt.Println(err11,"----")
	fmt.Println(map1)
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":map1})
}
func Test3(){

}

type Response struct {
	UserList []models.User`json:"UserList"`
	Page string `json:"Page"`
	Color string `json:"color"`
}



