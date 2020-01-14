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
	u.Nickname = "游客_" + u.Username
	result := u.InsertUser()
	if result == 1{
		c.JSON(http.StatusOK, gin.H{"Code": "1","Data":"注册成功"})
	}else if result == 0{
		c.JSON(http.StatusOK, gin.H{"Code": "0","Msg":"注册失败,存在相同用户名"})
	}else if result == 2{
		c.JSON(http.StatusOK, gin.H{"Code": "0","Msg":"注册失败"})
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
		utils.Error(fmt.Sprintf("%v","---------token创建失败------------",token_err))
		c.JSON(http.StatusOK, gin.H{"Code": "0","Msg":"token创建失败"})
	}
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":token})
}

func GetUserInfo(c *gin.Context){
	userInfo,err := utils.GetUser(c)
	if err != nil{
		utils.Error(fmt.Sprintf("------------用户信息获取失败-----------",err))
	}
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":userInfo})
}

func SetUserInfo(c *gin.Context){
	nickname := c.Query("Name")
	address := c.Query("Address")
	if len(nickname) == 0 || len(address) == 0{
		c.JSON(http.StatusOK, gin.H{"Code": "0","Msg":"地址和名字不能为空"})
	}
	userJson,ok := c.Get("userInfo")
	if !ok  {
		utils.Error(fmt.Sprintf("------------c.Get失败-----------"))
	}
	var user models.User
	err := json.Unmarshal(userJson.([] byte),&user)
	if err !=nil{
		fmt.Println("-----解析失败------")
	}
	user.Nickname = nickname
	user.Address = address
	result := user.UpdateUserNameAndAddress()
	c.JSON(http.StatusOK, gin.H{"Code": "200","Msg":result})
}

func Test2(c *gin.Context){
	var user []models.User
	global.GormConfig.Find(&user)
	mapresult := make(map[string]interface{})
	mapresult["userList"] = user
	mapresult["color"] = "red"
	err := utils.SetHashRedis("zzz","aa",mapresult)
	if err != nil{
		utils.Error(fmt.Sprintf("%v","-----------redis_set---------",err))
	}

	result,err1 := utils.GetHashRedis("zzz","aa")
	if err1 != nil{
		utils.Error(fmt.Sprintf("%v","-----------redis_get---------",err))
	}
	c.JSON(http.StatusOK, gin.H{"Code": "0","Data":result})
}
func Test3(){

}

type Response struct {
	UserList []models.User`json:"UserList"`
	Page string `json:"Page"`
	Color string `json:"color"`
}



