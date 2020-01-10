package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Query("token")
		if len(token) > 0{
			c.Next()
		}else {
			c.JSON(http.StatusOK,gin.H{
				"Code":333,
				"Msg":"token no",
			})
		}
		//c.Next()
	}
}

