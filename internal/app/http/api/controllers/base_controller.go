package controllers

import "github.com/gin-gonic/gin"

type BaseController struct{}

func (*BaseController) response(c *gin.Context, data *map[string]interface{}) {
	c.JSON(200, data)
}
