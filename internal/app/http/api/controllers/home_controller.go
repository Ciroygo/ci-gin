package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
	BaseController
}

// Home 首页
func (h *HomeController) Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "home page",
	})
}

// About 关于页面
func (h *HomeController) About(c *gin.Context) {
	h.response(c, &map[string]interface{}{"message": "ciroygo about ciroy"})
}

// GetName 路由中获取参数例子
func (h *HomeController) GetName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
