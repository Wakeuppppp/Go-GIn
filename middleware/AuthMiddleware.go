/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 22:13
 * @Project_Name : go-gin
 * @File : AuthMidlleware.go
 * @Software :GoLand
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin/common"
	"go-gin/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token格式错误"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token不合法"})
			ctx.Abort()
			return
		}

		// 验证通过，获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "用户不存在"})
			ctx.Abort()
			return
		}

		// 用户存在,将用户添加到上下文中
		ctx.Set("user", user)
		ctx.Next()
	}
}
