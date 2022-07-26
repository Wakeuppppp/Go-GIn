/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/27 1:00
 * @Project_Name : go-gin
 * @File : user_dto.go
 * @Software :GoLand
 */

package dto

import "go-gin/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
