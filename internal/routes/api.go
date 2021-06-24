package routes

import (
	"cigin/internal/app/http/api/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes 注册api相关路由
func RegisterApiRoutes(r *gin.Engine) {

	v1 := r.Group("/api")
	{
		hc := new(controllers.HomeController)
		v1.GET("/", hc.Home)
		v1.GET("/home", hc.Home)
		v1.GET("/about", hc.About)
		v1.GET("/getname/:name", hc.GetName)

		categories := new(controllers.CategoriesController)
		v1.GET("/categories", categories.Index)
	}
}
