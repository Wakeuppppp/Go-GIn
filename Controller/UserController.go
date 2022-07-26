/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 19:37
 * @Project_Name : go-gin
 * @File : UserController.go
 * @Software :GoLand
 */

package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/common"
	"go-gin/model"
	"go-gin/util"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"unicode"
)

func Login(ctx *gin.Context) {

}

func Register(ctx *gin.Context) {
	DB := common.GetDb()
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
		name = util.RandomString(10)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": name})
	}
	log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "注册成功",
		"name":      name,
		"telephone": telephone,
		"password":  password,
	})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
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
