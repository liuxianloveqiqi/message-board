package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/web_git/dao"
	"redrock/web_git/global"
	"redrock/web_git/service"
	"strconv"
)

// 登上界面
func Start(c *gin.Context) {
	c.SetCookie("login", "yes", 60, "/", "localhost", false, true)
	// 返回信息
	c.String(200, "Login success!")
}

// 注册
func Register(c *gin.Context) {
	//获取数据
	ri := c.PostForm("ID")
	global.RegisterID, _ = strconv.Atoi(ri)
	global.RegisterName = c.PostForm("name")
	global.RegisterPassword = c.PostForm("password")
	global.RegisterSecretProtection = c.PostForm("secretProtection")
	dao.InsertData(global.RegisterID, global.RegisterName, global.RegisterPassword, global.RegisterSecretProtection)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

// 登录
func Login(c *gin.Context) {
	//获取参数
	global.LoginName = c.PostForm("name")
	global.LoginPassword = c.PostForm("password")
	//验证(只需要用户名和密码)
	is := dao.QueryManyData(global.LoginName, global.LoginPassword)
	if is {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "登录成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "登录失败",
		})
		return
	}

}

// 密保查询
func SecretQurry(c *gin.Context) {
	li := c.PostForm("ID")
	global.LoginID, _ = strconv.Atoi(li)

	global.LoginSecretProtection = c.PostForm("secretProtection")
	//使用密保
	na, pa := dao.QueryRowData(global.LoginID, global.LoginSecretProtection)
	c.JSON(http.StatusOK, gin.H{
		"你的用户名为：": na,
		"你的密码为：":  pa,
	})
}

// 修改密码
func ResetPassword(c *gin.Context) {
	newPassword := c.PostForm("new-password")
	protationSecret := c.PostForm("serect")
	service.ResetPassword(newPassword, protationSecret, c)
}
