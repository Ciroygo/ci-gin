package configs

import "fmt"

// Initialize 配置信息初始化
func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法
	fmt.Printf("Initialize()配置信息加载完成\n")
}
