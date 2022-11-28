package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/web_git/global"
	"redrock/web_git/model"
)

// 注册数据
func InsertData(i int, n string, p string, ps string) {
	sqlStr := "insert into userMassage(ID,username,password,secretProtection) values (?,?,?,?)"
	r, err := global.DB.Exec(sqlStr, i, n, p, ps)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	ID, err2 := r.LastInsertId()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	fmt.Printf("ID: %v\n", ID)
}

// 查询密保
func QueryRowData(i int, s string) (string, string) {
	var n, p string
	sqlStr := "select ID,secretProtection from userMassage where ID=? and secretProtection=? "
	var u model.User
	err := global.DB.QueryRow(sqlStr, i, s).Scan(&u.Name, &u.Password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	n = u.Name
	p = u.Password
	return n, p
}

// 用户名和密码登录查询
func QueryManyData(n string, p string) bool {
	is := false
	sqlStr := "select username,password from usermassage"
	r, err := global.DB.Query(sqlStr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer r.Close()
	// 循环读取结果集中的数据
	for r.Next() {
		var u2 model.User
		err2 := r.Scan(&u2.Name, &u2.Password)
		if err2 != nil {
			fmt.Printf("err: %v\n", err2)
		}
		if u2.Name == n && u2.Password == p {
			is = true
			break
		}
	}
	return is
}

// 修改密码
func NewPassword(newpassword string, secret string, c *gin.Context) {
	strSql := "update usermessage set password=? where secretProtection=?"
	r, err := global.DB.Exec(strSql, newpassword, secret)
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
	err3 := global.DB.QueryRow("select password from usermessage where id=?", global.LoginID).Scan(&p)
	if err3 != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":  "200",
		"你的新密码为": p,
	})
}
