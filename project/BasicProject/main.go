package main

import (
	"BasicProject/dao/mysql"
	"BasicProject/router"
	"BasicProject/setting"
	"log"
)

func main() {
	if err := setting.Init(); err != nil {
		log.Println("setting Init ERROR:", err)
	}

	// 初始化MySQL
	if err := mysql.Init(setting.Conf.Mysql); err != nil {
		log.Println("init Mysql DB error")
	}

	// 初始化路由
	r := router.SetupRouter(setting.Conf.Mode)
	r.Run(":8085")

}
