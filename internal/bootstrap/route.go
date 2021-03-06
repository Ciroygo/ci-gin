package bootstrap

import (
	"cigin/internal/routes"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	router := gin.Default()
	routes.RegisterApiRoutes(router)
	routes.RegisterWebRoutes(router)

	return router
}
