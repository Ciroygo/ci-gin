package configs

import "fmt"

func init() {
	// 触发加载本目录下其他文件中的 init 方法
	fmt.Printf("app配置加载完成\n")
}
