package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_demo/models"
	"net/http"
)

func CheckUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Query("token")
		info,err := models.ParseToken(token)
		if err != nil{
			c.JSON(http.StatusOK,gin.H{
				"Code":333,
				"Msg":"token is invalid",
			})
			c.Abort()
			return
		}
		fmt.Println(info)
		userInfo := models.FindUserByUsername(info.Username)
		data,json_err := json.Marshal(userInfo)
		if json_err!=nil{
			fmt.Println("worryjson")
			return
		}
		//data := utils.ModelToMap(userInfo)
		c.Set("userInfo",data)
		c.Next()
	}
}

