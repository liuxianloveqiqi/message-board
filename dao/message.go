package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/web_git/global"
	"redrock/web_git/model"
)

func LeaveWords() {
	sqlStr := "insert into messageboard(senderID,sendername,receiveID,receivename,message) values (?,?,?,?,?)"
	r, err := global.DB.Exec(sqlStr, global.LoginID, global.LoginName, global.PostID, global.PostUerName, global.Message)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	i2, err2 := r.LastInsertId()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	fmt.Printf("i2: %v\n", i2)
}
func SeeWords(c *gin.Context) {
	sqlStr := "select senderID,sendername,message from messageboard where receivename=?"
	fmt.Println(global.LoginName)
	r, err := global.DB.Query(sqlStr, global.LoginName)
	fmt.Println(r)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer r.Close()
	// 循环读取结果集中的数据
	for r.Next() {
		err2 := r.Scan(&global.M.ID, &global.M.Name, &global.M.Messgae)
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
			return
		}
		if global.Message != "" {

			c.JSON(http.StatusOK, gin.H{
				"发送者ID":   global.M.ID,
				"发送者Name": global.M.Name,
				"发来的消息":   global.M.Messgae,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"发送者ID":   global.M.ID,
				"发送者Name": global.M.Name,
				"发来的消息":   "该留言已被删除",
			})
		}
	}
}
func SendComment() {
	strSql := "update messageboard set commenterID=?,commenterName=?,commend=? where senderid=? and receiveID=?"
	r, err := global.DB.Exec(strSql, global.LoginID, global.LoginName, global.CommendWords, global.SenderID, global.ReceiverID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	i2, err2 := r.RowsAffected()
	if err != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	fmt.Printf("i2: %v\n", i2)
}
func ShowComment(sendID int, receiveID int, c *gin.Context) {
	sqlStr := "select commenterID,ccommentName,comment from messageboard where senderID=? and  receiveID=?"

	r, err := global.DB.Query(sqlStr, sendID, receiveID)
	fmt.Println(r)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var m1 model.Message
	for r.Next() {
		err2 := r.Scan(&m1.CommentderID, &m1.CommentderName, &m1.Comment)
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
			return
		}
		if m1.Comment != "" {
			c.JSON(http.StatusOK, gin.H{
				"评论者ID":   m1.CommentderID,
				"评论者Name": m1.CommentderName,
				"评论的内容":   m1.Comment,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"评论者ID":   m1.CommentderID,
				"评论者Name": m1.CommentderName,
				"评论的内容":   "该评论已被删除",
			})
		}
	}
}
func NewMessage(si int, m string, c *gin.Context) {
	strSql := "update messageboard set message=? where receiveID=? and senderID=?"
	r, err := global.DB.Exec(strSql, m, si, global.LoginID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	i2, err2 := r.RowsAffected()
	if err != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	fmt.Printf("i2: %v\n", i2)
	var p string
	err3 := global.DB.QueryRow("select message from messageboard where senderID=? and receiveID=?", global.LoginID, si).Scan(&p)
	if err3 != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	if p != "" {
		c.JSON(http.StatusOK, gin.H{
			"state":     "200",
			"你的留言对象ID为": si,
			"你的新留言为":    p,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state": "200",
			"删除留言":  "成功",
		})
	}
}
func NewComment(si int, ri int, nc string, c *gin.Context) {
	strSql := "update messageboard set comment=? where receiveID=? and senderID=? and commenterID=?"
	r, err := global.DB.Exec(strSql, nc, si, ri, global.LoginID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	i2, err2 := r.RowsAffected()
	if err != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	fmt.Printf("i2: %v\n", i2)
	//查看新留言
	strsql2 := "select comment form messageboard where receiveID=? and senderID=? and commenterID=?"
	var p string
	err3 := global.DB.QueryRow(strsql2, ri, si, global.LoginID).Scan(&p)
	if err3 != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	if p != "" {
		c.JSON(http.StatusOK, gin.H{
			"state":         "200",
			"你评论的留言的发送者ID为": si,
			"你评论的留言的接受者ID为": ri,
			"你的评论修改为":       p,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state":         "200",
			"你评论的留言的发送者ID为": si,
			"你评论的留言的接受者ID为": ri,
			"删除评论":          "成功",
		})
	}
}
