package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	us "sxk.go-kit/internal/user/service"
)

/*
@Time    : 2021/3/9 18:41
@Author  : austsxk
@Email   : austsxk@163.com
@File    : userEndpoint.go
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

// 生成端点,时需要进行业务逻辑处理，也就是要将具体实现对象注入到内部，然后返回endpoint对象接即可
func GenUserEndpoint(u us.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// 断言请求为UserRequest
		r := request.(UserRequest)
		// 调用接口获取响应
		data := u.SearchUser(r.ID)
		return UserResponse{Data: data}, nil

	}
}
