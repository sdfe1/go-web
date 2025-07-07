package main

import (
	"Project1/config"
	"Project1/model"
	"Project1/router"
	"fmt"
)

func main() {
	fmt.Printf("开始调用")
	//1.加载配置
	config.InitConfig()
	//2.初始化数据库
	model.InitDb()
	//3.设置路由
	router.InitRouter()
}
