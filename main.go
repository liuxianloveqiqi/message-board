package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"redrock/web_git/cmd"
)

func main() {
	r := gin.Default()
	cmd.Userroute(r)
	//cmd.Messageroute(r)
	r.Run()
}
