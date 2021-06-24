package controllers

import (
	"cigin/internal/app/models/category"
	"cigin/pkg/model"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseController
}

// Index 分类列表
func (h *CategoriesController) Index(c *gin.Context) {
	var categories []category.Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return
	}
	h.response(c, categories)
}

// Show 分类详情
func (h *CategoriesController) Show(c *gin.Context) {
	var categories []category.Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return
	}
	c.JSON(200, categories)
}

// Store 添加分类
func (h *CategoriesController) Store(c *gin.Context) {
	var categories []category.Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return
	}
	c.JSON(200, categories)
}

// Store 修改分类
func (h *CategoriesController) Update(c *gin.Context) {
	var categories []category.Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return
	}
	c.JSON(200, categories)
}

// Delete 删除分类
func (h *CategoriesController) Delete(c *gin.Context) {
	var categories []category.Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return
	}
	c.JSON(200, categories)
}
