package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/web_git/dao"
	"redrock/web_git/global"
	"redrock/web_git/service"
	"strconv"
)

// 使用中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("login"); err == nil {
			if cookie == "yes" {
				c.Next()
			}
		} else {
			// 返回错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "没有登录"})
			// 若验证不通过，不再调用后续的函数处理
			c.Abort()
		}
	}
}

// 发送留言
func SendWords(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"状态": 200,
	})
	p := c.PostForm("postID")
	global.PostID, _ = strconv.Atoi(p)
	global.PostUerName = c.PostForm("postname")
	global.Message = c.PostForm("message")
	fmt.Println("------------------------")
	fmt.Println(global.PostID, global.PostUerName, global.Message)
	dao.LeaveWords()
}

// 查看留言
func LookWords(c *gin.Context) {
	dao.SeeWords(c)
}

// 根据留言发送者ID和接收者ID 进行留言评价
func Commend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"状态": 200,
	})
	sid := c.PostForm("senderID")
	global.SenderID, _ = strconv.Atoi(sid)
	rid := c.PostForm("receiverID")
	global.ReceiverID, _ = strconv.Atoi(rid)
	global.CommendWords = c.PostForm("commend")
	dao.SendComment()
	c.JSON(http.StatusOK, gin.H{
		"评价成功": "OK",
	})
}

// 根据发送者和接受者ID 查看一条留言下的所有评论
func LookCommend(c *gin.Context) {
	sid := c.PostForm("senderID")
	sendID, _ := strconv.Atoi(sid)
	rid := c.PostForm("receiverID")
	receiveID, _ := strconv.Atoi(rid)
	service.LookCmommend(sendID, receiveID, c)
}

// 根据接收者ID修改留言
func Retmessage(c *gin.Context) {
	newmessage := c.PostForm("new-message")
	rid := c.PostForm("receiverID")
	receiveID, _ := strconv.Atoi(rid)
	service.RetMessage(receiveID, newmessage, c)
}

// 根据接收者ID删除留言
func DelectMessage(c *gin.Context) {
	rid := c.PostForm("receiverID")
	receiveID, _ := strconv.Atoi(rid)
	service.RetMessage(receiveID, "", c)
}

// 根据发送者ID和接受者ID修改评论
func RetComment(c *gin.Context) {
	sid := c.PostForm("senderID")
	sendID, _ := strconv.Atoi(sid)
	rid := c.PostForm("receiverID")
	receiveID, _ := strconv.Atoi(rid)
	newcomment := c.PostForm("new-comment")
	service.RetCommend(sendID, receiveID, newcomment, c)
}

// 根据发送者ID和接受者ID删除评论
func DelectComment(c *gin.Context) {
	sid := c.PostForm("senderID")
	sendID, _ := strconv.Atoi(sid)
	rid := c.PostForm("receiverID")
	receiveID, _ := strconv.Atoi(rid)
	service.RetCommend(sendID, receiveID, "", c)
}
