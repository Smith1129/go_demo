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
	data,err := json.Marshal(user)
	if err !=nil{
		fmt.Println(err)
	}
	mapresult := map[string]string{
		"userList":string(data),
		"page":"1",
		"color":"red",
	}
	//var mapresult map[string]string
	//mapresult["userList"] = string(data)
	//mapresult["page"] = "1"
	//mapresult["color"] = "red"
	fmt.Println(mapresult["page"])
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":mapresult})
	//fmt.Println(string(data))
	//key := "userList"
	//result,err := redisClient.HSet(key,"aa",mapresult).Result()
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(result,err)
	//fmt.Println("--------",result)
	//hdata := redisClient.HGetAll(key).Val()
	//fmt.Println(hdata)
	//Test3()



	//string操作
	//value := "zzz"
	//key := "name"
	//redisClient.Set(key,value,0)
	//val := redisClient.Get(key)
	//if val == nil {
	//	fmt.Errorf("StringDemo get error")
	//}
	//fmt.Println("value---------",val.Val())


	//list操作
	//listkey := "list"
	//_,err := redisClient.RPush(listkey,"a","b","c").Result()
	//if err!=nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(redisClient.LRange(listkey,0,2))
}
func Test3(){
	key := "userList"
	redisClient := global.GetRedisClient()
	val := redisClient.Get(key)
	if val == nil{
		fmt.Println("worry")
	}
	fmt.Println(val.Val())
}

type Response struct {
	UserList []models.User`json:"UserList"`
	Page string `json:"Page"`
	Color string `json:"color"`
}


