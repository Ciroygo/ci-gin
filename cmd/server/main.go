package main

import (
	"cigin/internal/bootstrap"

	"cigin/configs"
)

func init() {
	// config 初始化
	configs.Initialize()
	// 数据初始化
	bootstrap.SetupDB()
}

func main() {
	// 路由注册
	router := bootstrap.SetupRoute()

	router.Run(":8091")
}
