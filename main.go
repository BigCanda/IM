package main

import (
	"Im/router"
	"Im/utils"
)

func main() {
	// 初始化设置
	utils.InitConfig()
	// 初始化Mysql设置
	utils.InitMysql()
	// 调用控制器
	r := router.Router()
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") 可以通过修改r.Run(":xxxx")修改端口号
	if err != nil {
		panic(err)
	}
}
