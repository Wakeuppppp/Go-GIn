/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2022/7/30 14:28
 * @Project_Name : go-gin
 * @File : category_request.go
 * @Software :GoLand
 */

package vo

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required" `
}
