package service

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

/*
@Time    : 2021/3/11 16:43
@Author  : austsxk
@Email   : austsxk@163.com
@File    : user_service.go
@Software: GoLand
*/

// 对用户的请求进行序列化， 与服务端相反
func EncodeUserRequest(ctx context.Context, request *http.Request, data interface{}) error {
	d := data.(UserRequest)
	request.URL.Path += "/user/" + strconv.Itoa(d.ID)
	return nil
}

// 对用户的响应进行解码, 与服务端相反
func DecodeUserResponse(ctx context.Context, res *http.Response) (response interface{}, err error) {
	// 将响应断言为定义的响应
	if res.StatusCode >= 400 {
		return nil, errors.New("data not found")
	}
	// 定义一个响应对象，进行存储解析的响应
	var userResponse UserResponse
	err = json.NewDecoder(res.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, err
}

// add user encode
func EncodeAddUserRequest(ctx context.Context, request *http.Request, data interface{}) error {
	d := data.(AddUserRequest)
	byteDate, err := json.Marshal(d)
	if err != nil {
		return err
	}
	request.URL.Path += "/user/add"
	// data set in body
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Body = ioutil.NopCloser(strings.NewReader(string(byteDate[:])))
	return nil
}

func DecodeAddUserResponse(ctx context.Context, res *http.Response) (response interface{}, err error) {
	// 将响应断言为定义的响应
	if res.StatusCode >= 400 {
		return nil, errors.New("data not found")
	}
	// 定义一个响应对象，进行存储解析的响应
	var r AddUserResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return r, err
}
