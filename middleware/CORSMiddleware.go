/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/28 11:17
 * @Project_Name : go-gin
 * @File : CORSMiddleware.go
 * @Software :GoLand
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")   // 最大缓时间
		ctx.Writer.Header().Set("Access-Control-Allow-Method", "*")  // 允许通过的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*") //
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
