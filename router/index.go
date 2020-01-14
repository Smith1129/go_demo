package router

import (
	"github.com/gin-gonic/gin"
	"go_demo/middlewares"
	"go_demo/router/api"
	"io"
	"os"
)


func RouterConfig() *gin.Engine {
	router := gin.New()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router.GET("/api/register",api.Register)
	router.GET("/api/login",api.Login)
	router.GET("/api/goodlist",api.GetGoodList)
	//router.GET("/test",api.Test)
	router.GET("/test2",api.Test2)
	apiGroup := router.Group("/api")
	apiGroup.Use(middlewares.CheckUser())
	{
		apiGroup.GET("/userInfo",api.GetUserInfo)
		apiGroup.GET("/userInfoSet",api.SetUserInfo)
	}
	return router
}
