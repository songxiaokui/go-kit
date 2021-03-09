package transport

import (
	"context"
	"encoding/json"
	"errors"
	mymux "github.com/gorilla/mux"
	"net/http"
	"strconv"
	ue "sxk.go-kit/internal/user/endpoint"
)

/*
@Time    : 2021/3/9 18:58
@Author  : austsxk
@Email   : austsxk@163.com
@File    : userTransport.go
@Software: GoLand
*/

// 用来解析请求和返回响应，可以自定义http handler

// 请求处理
func EncodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	if data, ok := mymux.Vars(r)["id"]; ok {
		d, _ := strconv.Atoi(data)
		return ue.UserRequest{ID: d}, nil
	}
	return nil, errors.New("NotFound")
}

// 响应处理
func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// 设置响应头信息
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
