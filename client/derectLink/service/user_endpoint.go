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

// 定义新增用户接口，使用进程内部缓存技术
type AddUserRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AddUserResponse struct {
	Status bool `json:"status"`
}
