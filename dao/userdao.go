package dao

import (
	"fmt"
	"redrock/web_git/model"
)

// 注册数据
func InsertData(i int, n string, p string, ps string) {
	sqlStr := "insert into userMassage(ID,username,password,secretProtection) values (?,?,?,?)"
	r, err := db.Exec(sqlStr, i, n, p, ps)
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
	err := db.QueryRow(sqlStr, i, s).Scan(&u.Name, &u.Password)
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
	r, err := db.Query(sqlStr)
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
