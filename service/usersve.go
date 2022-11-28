package service

import (
	"github.com/gin-gonic/gin"
	"redrock/web_git/dao"
)

// 修改密码
func ResetPassword(np string, ps string, c *gin.Context) {
	dao.NewPassword(np, ps, c)

}
