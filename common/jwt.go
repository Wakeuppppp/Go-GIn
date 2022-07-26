/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 20:50
 * @Project_Name : go-gin
 * @File : jwt.go
 * @Software :GoLand
 */

package common

import (
	"github.com/dgrijalva/jwt-go"
	"go-gin/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "谷安-Gin-Vue",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

/*
token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
		eyJVc2VySWQiOjMsImV4cCI6MTY1OTQ0ODk5NCwiaWF0IjoxNjU4ODQ0MTk0LCJpc3MiOiLosLflroktR2luLVZ1ZSIsInN1YiI6InVzZXIgdG9rZW4ifQ.
		_o6az8FyIvlJkSt3XppSL1y_kYmHUZRi6P7qPYX5vI4"
第一部分是协议头，储存的是token使用的加密协议
第二部分是负载，储存的是claims结构体信息
第三部分是 前面两部分+jwtKey 哈希出来的值

可以用 "echo 任一部分的字符串 | base64 -d" 解析出结果
比如解析第二部分
"echo eyJVc2VySWQiOjMsImV4cCI6MTY1OTQ0ODk5NCwiaWF0IjoxNjU4ODQ0MTk0LCJpc3MiOiLosLflroktR2luLVZ1ZSIsInN1YiI6InVzZXIgdG9rZW4ifQ | base64 -d"
结果：{"UserId":3,"exp":1659448994,"iat":1658844194,"iss":"谷安-Gin-Vue","sub":"user token"}
跟claims结构体一一对应

第一部分：{"alg":"HS256","typ":"JWT"}

*/

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
