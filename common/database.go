/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 19:46
 * @Project_Name : go-gin
 * @File : database.go
 * @Software :GoLand
 */

package common

import (
	"fmt"
	"github.com/spf13/viper"
	"go-gin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{}) // 自动按照User格式创建一个user表
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
