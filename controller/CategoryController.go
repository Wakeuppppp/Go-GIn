/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/29 20:06
 * @Project_Name : go-gin
 * @File : CategoryController.go
 * @Software :GoLand
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin/common"
	"go-gin/model"
	"go-gin/response"
	"go-gin/vo"
	"gorm.io/gorm"
	"strconv"
)

type ICateGoryController interface {
	RestController
}

type CateGoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICateGoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CateGoryController{db}
}

func (c CateGoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必须填写")
		return
	}

	category := model.Category{Name: requestCategory.Name}

	c.DB.Create(&category)
	response.Success(ctx, gin.H{"category:": category}, "")
}

func (c CateGoryController) Update(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "数据验证错误，分类名称必须填写")
		return
	}
	category := model.Category{Name: requestCategory.Name}
	id, _ := strconv.Atoi(ctx.Param("id"))

	var updateCategory model.Category
	c.DB.Where("ID = ?", id).First(&updateCategory)
	if updateCategory.ID == 0 {
		response.Fail(ctx, nil, "数据未找到")
		return
	}
	c.DB.Model(&updateCategory).Update("name", category.Name)
	response.Success(ctx, gin.H{"updateCategory": category}, "修改成功")
}

func (c CateGoryController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var showCategory model.Category
	c.DB.Where("ID = ?", id).First(&showCategory)
	if showCategory.ID == 0 {
		response.Fail(ctx, nil, "数据未找到")
		return
	}
	response.Success(ctx, gin.H{"showCategory": showCategory}, "")
}

func (c CateGoryController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// if err := c.DB.Delete(model.Category{}, id).Error; err != nil {
	// 	response.Fail(ctx, nil, "删除失败，请重试")
	// 	return
	// }

	var deleteCategory model.Category
	c.DB.Where("ID = ?", id).First(&deleteCategory)
	if deleteCategory.ID == 0 {
		response.Fail(ctx, nil, "数据未找到")
		return
	}
	c.DB.Delete(&deleteCategory, id)
	response.Success(ctx, nil, "删除成功")
}
