/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 19:52
 * @Project_Name : go-gin
 * @File : routes.go
 * @Software :GoLand
 */

package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/Controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", Controller.Register)
	return r
}
