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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"
	"unicode"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	fmt.Println("连接数据库")
	db := InitDB()
	fmt.Println("连接成功")

	fmt.Println(db)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		// 数据验证
		if !CheckPhone(telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号格式错误"})
			return
		}
		if !CheckPwd(password) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码需包含大小写、数字、特殊字符，长度在8~20位"})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": name})
		}
		log.Println(name, telephone, password)

		// 判断手机号是否存在
		if isTelephoneExist(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
			return
		}

		// 创建用户
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		// 返回结果
		ctx.JSON(http.StatusOK, gin.H{
			"msg":       "注册成功",
			"name":      name,
			"telephone": telephone,
			"password":  password,
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// CheckPhone 检验手机号
func CheckPhone(phone string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	return reg.MatchString(phone)
}

// CheckPwd 检验密码
func CheckPwd(password string) bool {
	if len(password) < 8 || len(password) > 20 {
		fmt.Println("长度不对")
		return false
	}
	mask := 0
	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			mask |= 1
		case unicode.IsUpper(c):
			mask |= 2
		case unicode.IsDigit(c):
			mask |= 4
		default:
			mask |= 8
		}
	}
	return mask == 15
}

// RandomString 随机生成10位数的用户名
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "gin-vue"
	username := "root"
	password := "root"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&User{}) // 自动按照User格式创建一个user表
	return db
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
