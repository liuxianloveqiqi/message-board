package dao

import (
	"database/sql"
	"fmt"
	"redrock/web_git/global"
)

func Opendata() {
	//连接数据库
	err := InitDB()
	if err != nil {
		fmt.Printf("------err--------: %v\n", err)
	} else {
		fmt.Println("---------连接成功------------")
	}
	fmt.Printf("db: %v\n", global.DB)
}
func InitDB() (err error) {
	dsn := "root:xian712525@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4"
	// open函数只是验证格式是否正确，并不是创建数据库连接
	global.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 与数据库建立连接
	err2 := global.DB.Ping()
	if err2 != nil {
		return err2
	}
	return nil
}
