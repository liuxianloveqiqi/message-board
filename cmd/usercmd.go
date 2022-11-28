package cmd

import (
	"github.com/gin-gonic/gin"
	"redrock/web_git/api"
	"redrock/web_git/dao"
)

func Userroute(r *gin.Engine) {
	//用户数据库准备
	dao.Opendata()
	//用户路由准备
	us := r.Group("/user")
	{
		us.POST("/register", api.Register)           //注册
		us.GET("/start", api.Start)                  //登上登录界面
		us.POST("/login", api.Login)                 //登录
		us.POST("/secret", api.SecretQurry)          //密保
		us.POST("/resetpassword", api.ResetPassword) //修改密码
	}
}
