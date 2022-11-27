package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/web_git/dao"
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
	registerID, _ := strconv.Atoi(ri)
	registerName := c.PostForm("name")
	registerPassword := c.PostForm("password")
	registerSecretProtection := c.PostForm("secretProtection")
	dao.InsertData(registerID, registerName, registerPassword, registerSecretProtection)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

// 登录
func Login(c *gin.Context) {
	//获取参数
	loginName := c.PostForm("name")
	loginPassword := c.PostForm("password")
	//验证(只需要用户名和密码)
	is := dao.QueryManyData(loginName, loginPassword)
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
	loginID, _ := strconv.Atoi(li)
	loginSecretProtection := c.PostForm("secretProtection")
	//使用密保
	na, pa := dao.QueryRowData(loginID, loginSecretProtection)
	c.JSON(http.StatusOK, gin.H{
		"你的用户名为：": na,
		"你的密码为：":  pa,
	})
}
