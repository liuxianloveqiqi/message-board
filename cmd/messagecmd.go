package cmd

import (
	"github.com/gin-gonic/gin"
	"redrock/web_git/api"
)

func Messageroute(r *gin.Engine) {
	ms := r.Group("/message", api.AuthMiddleWare())
	{
		ms.POST("/sendword", api.SendWords)          //发送留言
		ms.POST("/lookword", api.LookWords)          //查看留言
		ms.POST("/comment", api.Commend)             //进行留言评价
		ms.POST("/lookcomment", api.LookCommend)     //查看一条留言下的所有评论
		ms.POST("/retmessage", api.Retmessage)       //修改留言
		ms.POST("/delectmessage", api.DelectMessage) //删除留言
		ms.POST("/retcomment", api.RetComment)       //修改评论
		ms.POST("/delectcomment", api.DelectComment) //删除评论

	}
}
