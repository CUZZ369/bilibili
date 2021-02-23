package cmd

import (
	"github.com/gin-gonic/gin"
	"test01/controller"
	"test01/middleware"
)
func Entrance() {
	router:=gin.Default()
	//不需要登录的部分
    v1:=router.Group("/BiliBili/user")
    {
		v1.POST("/register",controller.Register)
		v1.GET("/login",controller.Login)
		v1.GET("/showUser",controller.ShowUser)//查看个人页面
	}
   	//需要登录的部分
    v1.Use(middleware.Validate())
    {   //视频
    	V2:=v1.Group("/Video")
    	{
			V2.POST("/makeVideo",controller.CreateVideo)
			V2.GET("/VideoDetail",controller.ShowVideo)//查看视频详情
		}
	}
	router.Run(":8080")
}