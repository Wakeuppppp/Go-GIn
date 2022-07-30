/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/29 20:10
 * @Project_Name : go-gin
 * @File : RestController.go
 * @Software :GoLand
 */

package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
}
