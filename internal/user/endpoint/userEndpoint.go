package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"strconv"
	div "sxk.go-kit/api/discovery/user"
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
		// 调用接口获取响应, 为了区别不同的服务端口，加上服务端口标示
		data := u.SearchUser(r.ID) + " - " + strconv.Itoa(div.ServicePort)
		return UserResponse{Data: data}, nil

	}
}

// 定义新增用户接口，使用进程内部缓存技术
type AddUserRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AddUserResponse struct {
	Status bool `json:"status"`
}

func GenAddUserEndpoint(u us.UserServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(AddUserRequest)
		result := u.AddUser(r.Name, r.ID)
		return AddUserResponse{Status: result}, nil
	}
}
