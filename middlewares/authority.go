package middlewares

import (
	"encoding/json"
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
		data,_ := json.Marshal(info)
		c.Set("userInfo",data)
		c.Next()
	}
}

