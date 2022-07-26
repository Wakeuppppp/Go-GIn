/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 19:43
 * @Project_Name : go-gin
 * @File : util.go
 * @Software :GoLand
 */

package util

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"unicode"
)

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

// CheckPhone 检验手机号是否合法
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

// CheckPwd 检验密码是否合法
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
