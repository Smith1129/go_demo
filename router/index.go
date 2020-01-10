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
	apiGroup := router.Group("/api")
	router.POST("/api/register",api.Register)
	router.GET("/test",api.Test)
	apiGroup.Use(middlewares.CheckUser())
	{
		apiGroup.GET("/test",api.Search)
	}

	//router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// 你的自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	//router.Use(gin.Recovery())

	return router
}
