package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test01/service"
)
//注册
func Register(ctx *gin.Context) {
	res := service.Register(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "注册成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "failure",
		})
	}
}
//登录
func Login(ctx *gin.Context) {
	res := service.Login(ctx)
	if res {
		ctx.SetCookie("username", ctx.PostForm("username"), 240, "/", "127.0.0.1", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "欢迎回来" + ctx.PostForm("username"),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "用户名或密码错误！",
		})
	}
}
//上传视频
func CreateVideo(c *gin.Context)  {
	service := service.CreateVideo(c)
	if service {
		c.JSON(http.StatusOK, gin.H{
			"code":    10000,
			"message": "视频上传成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    20000,
			"message": "视频上传失败",
		})

	}
}
//查看视频页详情
func ShowVideo(c *gin.Context)  {
	th := service.ShowVideo(c)
	err:=c.ShouldBind(&th)
	if err==nil{
		res := service.ShowVideo(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "查看失败",
		})

	}
}
func ShowUser(c *gin.Context)  {
	th := service.ShowUser(c)
	err:=c.ShouldBind(&th)
	if err==nil{
		res := service.ShowVideo(c)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "查看失败",
		})

	}
}