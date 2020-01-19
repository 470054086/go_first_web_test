package main

import (
	"first_web/bootstrap"
	_ "first_web/bootstrap/database"
	_ "first_web/bootstrap/ini"
	_ "first_web/bootstrap/log"
	_ "first_web/router"
)
func main()  {
	// 启动进程
	bootstrap.Func.Start()
}