package controllers

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (*BaseController) response(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}
