package routes

import (
	"cigin/internal/app/http/web/controllers"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes 注册api相关路由
func RegisterWebRoutes(r *gin.Engine) {
	// 加载前端模版文件
	r.HTMLRender = loadTemplates("internal/resources/views")

	fmt.Println(r.HTMLRender)
	// 加载静态资源
	r.Static("/assets", "web/assets")

	hc := new(controllers.PagesController)
	r.GET("/home", hc.Home)
}

// loadTemplates 前端模版文件加载
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	common, err := filepath.Glob(templatesDir + "/layouts/common/*.gohtml")

	if err != nil {
		panic(err.Error())
	}

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.gohtml")
	layouts = append(layouts, common...)

	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/**/*.gohtml")

	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	var tmp string
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		tmp = strings.Replace(include, "internal/resources/views/includes/", "", -1)
		tmp = strings.Replace(tmp, ".gohtml", "", -1)
		tmp = strings.Replace(tmp, "/", ".", -1)
		r.AddFromFiles(tmp, files...)
		fmt.Println(files)
		fmt.Println("页面：" + tmp)
	}
	return r
}
