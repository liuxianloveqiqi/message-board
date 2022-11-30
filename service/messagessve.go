package service

import (
	"github.com/gin-gonic/gin"
	"redrock/web_git/dao"
)

// 查看评论
func LookCmommend(si int, ri int, c *gin.Context) {

	dao.ShowComment(si, ri, c)
}

// 修改留言
func RetMessage(si int, m string, c *gin.Context) {

	dao.NewMessage(si, m, c)
}

// 修改评论
func RetCommend(si int, ri int, nc string, c *gin.Context) {

	dao.NewComment(si, ri, nc, c)
}
