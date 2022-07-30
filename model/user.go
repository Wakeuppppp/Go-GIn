/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/26 19:34
 * @Project_Name : go-gin
 * @File : user.go
 * @Software :GoLand
 */

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
	Note      string `gorm:"size:255;not null"`
}
