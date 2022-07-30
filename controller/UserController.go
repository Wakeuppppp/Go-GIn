/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 19:37
 * @Project_Name : go-gin
 * @File : UserController.go
 * @Software :GoLand
 */

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin/common"
	"go-gin/dto"
	"go-gin/model"
	"go-gin/response"
	"go-gin/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	// 使用gin框架提供的Bind函数
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}

	// 获取参数
	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在，请先注册"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发送token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Println("token generate error: ", err)
		return
	}

	// 返回结果
	// ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登陆成功"})
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 使用map获取请求参数
	// var requestUser = make(map[string]string)
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	// fmt.Println("map: ", requestUser)

	// 使用结构体来获取参数
	// var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	// fmt.Println("结构体: ", requestUser)

	// 使用gin框架提供的Bind函数
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}

	// 获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	// name := "requestUser.Name"
	// telephone := "requestUser.Telephone"
	// password := "requestUser.Password"

	// 数据验证
	fmt.Println("Gin: ", requestUser.Name, requestUser.Telephone, requestUser.Password)
	if !util.CheckPhone(telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号格式错误")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号格式错误"})
		return
	}
	if !util.CheckPwd(password) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码需包含大小写、数字、特殊字符，长度在8~20位")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码需包含大小写、数字、特殊字符，长度在8~20位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
		// response.Response(ctx,http.StatusUnprocessableEntity,422,nil,name)
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": name})
	}
	// log.Println(name, telephone, password)

	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		// ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
		Note:      password,
	}
	DB.Create(&newUser)
	fmt.Println("创建用户成功: ", newUser.Name, newUser.Telephone)

	// 发送token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Println("token generate error: ", err)
		return
	}

	// 返回结果
	// ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登陆成功"})
	response.Success(ctx, gin.H{"token": token}, "注册成功")

	// 返回结果
	// response.Success(ctx, nil, "注册成功")
	// ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}}) // user.(model.User)类型断言
}
