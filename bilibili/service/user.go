package service

import (
	"github.com/gin-gonic/gin"
	"test01/models"
)

func Register(ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	res := models.Register(username, password)
	return res
}

func Login(ctx *gin.Context) bool {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	res := models.Login(username, password)
	return res
}