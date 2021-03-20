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
	discovery "sxk.go-kit/api/discovery/user"
	. "sxk.go-kit/client/derectLink/service"
)

func main() {
	// parse url
	t, _ := url.Parse(fmt.Sprintf("http://%s:9004", discovery.DefaultAddress))

	// add user info

	clientPost := http.NewClient("POST", t, EncodeAddUserRequest, DecodeAddUserResponse)
	addPoint := clientPost.Endpoint()
	rsp, err := addPoint(context.Background(), AddUserRequest{ID: 2, Name: "austsxk"})

	d := rsp.(AddUserResponse)
	fmt.Printf("add status : %#v\n", d.Status)

	// get user info
	client := http.NewClient("GET", t, EncodeUserRequest, DecodeUserResponse)
	endPoint := client.Endpoint()
	response, err := endPoint(context.Background(), UserRequest{ID: 2})
	if err != nil {
		fmt.Println(err)
		return
	}
	userInfo := response.(UserResponse)
	fmt.Printf("recive response: %s\n", userInfo.Data)
}
