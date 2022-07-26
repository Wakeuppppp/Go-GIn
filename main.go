/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/24 22:39
 * @Project_Name : go-gin
 * @File : main.go.go
 * @Software :GoLand
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/common"
)

func main() {
	fmt.Println("连接数据库")
	db := common.InitDB()
	fmt.Println("连接成功")
	fmt.Println(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
