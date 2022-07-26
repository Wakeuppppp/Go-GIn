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
	"github.com/spf13/viper"
	"go-gin/common"
	"os"
)

func main() {
	fmt.Println("初始化配置")
	InitConfig()
	fmt.Println("初始化完成")
	fmt.Println("连接数据库")
	common.InitDB()
	fmt.Println("连接成功")

	r := gin.Default()
	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// config := viper.New()

	viper.AddConfigPath(path + "/config") // 设置读取的文件路径
	viper.SetConfigName("application")    // 设置读取的文件名
	viper.SetConfigType("yaml")           // 设置文件的类型
	// 尝试进行配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
