package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test01/models"
)
//身份验证中间件
func Validate() gin.HandlerFunc{
	return func(c *gin.Context){
		username:=c.Query("username")
		password:=c.Query("password")
		res :=models.Login(username,password)
		if res {
			c.JSON(http.StatusOK,gin.H{"message":"已登录"})
			c.Next()
		}else {
			c.JSON(http.StatusUnauthorized,gin.H{"message":"请先登录"})
		}
	}
}
