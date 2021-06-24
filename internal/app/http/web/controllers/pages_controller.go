package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PagesController struct {
}

// Home 首页
func (*PagesController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "pages.home", gin.H{
		"title": "Html5 Article Engine",
	})
}
