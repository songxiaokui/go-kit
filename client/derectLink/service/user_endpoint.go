package service

/*
@Time    : 2021/3/11 16:43
@Author  : austsxk
@Email   : austsxk@163.com
@File    : user_endpoint.go
@Software: GoLand
*/

// 定义请求时的请求对象
type UserRequest struct {
	ID int `json:"id"`
}

// 定义响应时的响应对象
type UserResponse struct {
	Data string `json:"data"`
}
