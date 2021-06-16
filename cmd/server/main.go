package main

import (
	"cigin/internal/bootstrap"

	"cigin/configs"
)

func init() {
	// config 注册
	configs.Initialize()
}

func main() {
	// 数据库注册

	// 路由注册
	router := bootstrap.SetupRoute()

	router.Run(":80")
}
