package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"test01/models"
)

func CreateVideo(ctx *gin.Context) bool {
	th := models.CreateVideoService{}
	th.Info = ctx.PostForm("title")
	th.Info = ctx.PostForm("info")
	th.Username= ctx.PostForm("username")
	res := models.CreateVideo(&th)
	return res
}
//根据id查询视频详情
func ShowVideo(ctx *gin.Context) models.VideoDetail {
	id:=ctx.PostForm("id")
	//title:=ctx.PostForm("title")
	//类型转化
	num,_:=strconv.Atoi(id)
	detail:=models.ShowVideoId(num)
	return detail
}

func ShowUser(ctx *gin.Context) models.UserDetail {
	Username:=ctx.PostForm("username")
	detail :=models.ShowUser(Username)
    return detail
}