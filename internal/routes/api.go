package routes

import (
	"cigin/internal/app/http/api/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes 注册api相关路由
func RegisterApiRoutes(r *gin.Engine) {
	hc := new(controllers.HomeController)
	r.GET("/ping", hc.Home)
	r.GET("/about", hc.About)
	r.GET("/getname/:name", hc.GetName)
}
