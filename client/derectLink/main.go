package main

/*
@Time    : 2021/3/11 16:41
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/transport/http"
	"net/url"
	. "sxk.go-kit/client/derectLink/service"
)

func main() {
	// parse url
	t, _ := url.Parse("http://192.168.31.102:9999")
	// get user info
	client := http.NewClient("GET", t, EncodeUserRequest, DecodeUserResponse)
	endPoint := client.Endpoint()
	response, err := endPoint(context.Background(), UserRequest{ID: 1})
	if err != nil {
		fmt.Println(err)
		return
	}
	userInfo := response.(UserResponse)
	fmt.Printf("recive response: %s\n", userInfo.Data)

	// add user info
	// clientPost := http.NewClient("POST", t, EncodeAddUserRequest, DecodeAddUserResponse)

}
