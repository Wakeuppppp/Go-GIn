/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/29 20:01
 * @Project_Name : go-gin
 * @File : category.go
 * @Software :GoLand
 */

package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`
}
