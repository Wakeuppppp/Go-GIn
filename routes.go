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
	"go-gin/controller"
	"go-gin/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	// r.GET("api/auth/info", controller.Info)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	return r
}
