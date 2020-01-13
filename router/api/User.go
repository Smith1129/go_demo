package api

import (
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
	fmt.Println(userInfo.ID)
	//userInfo,ok := c.Get("userInfo")
	//if !ok{
	//	return
	//}
	//fmt.Println(reflect.TypeOf(userInfo))
	//var mapResult map[string]interface{}
	//err := json.Unmarshal(userInfo.([] byte), &mapResult)
	//if err != nil {
	//	fmt.Println("JsonToMapDemo err: ", err)
	//	return
	//}
	//var user []models.User
	//user = userInfo
	//fmt.Println(mapResult["username"])

	//fmt.Println(reflect.ValueOf(userInfo["Username"]))
	//fmt.Println(a1.ID)
}

func Test2(c *gin.Context){
	var user []models.User
	global.GormConfig.Find(&user)
	m := make(map[string] string)
	m["ss"] = "dd"
	//res := &Response{
	//	UserList: user,
	//	Page:     "1",
	//	Color:    "2",
	//}
	//token := c.Query("token")
	//info,err := models.ParseToken(token)
	//if err != nil{
	//	fmt.Println(err)
	//}
	c.JSON(http.StatusOK, gin.H{"Code": "200","Data":m})
}

type Response struct {
	UserList []models.User`json:"UserList"`
	Page string `json:"Page"`
	Color string `json:"color"`
}


