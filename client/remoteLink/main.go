package main

/*
@Time    : 2021/3/12 14:15
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	transportHttp "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	"io"
	"net/url"
	"os"
	discovery "sxk.go-kit/api/discovery/user"
	derectTranspot "sxk.go-kit/client/derectLink/service"
	"time"
)

func main() {
	// 1. 创建consul客户端
	config := consulapi.DefaultConfig()
	config.Address = discovery.DefaultAddress + ":8500"
	// 创建consul client
	apiClient, _ := consulapi.NewClient(config)
	// 使用kit库下的consul创建客户端
	client := consul.NewClient(apiClient)

	// 使用kit库中的日志类型
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)

	// 2. 创建consul实例
	var Tag = []string{"userModel"}
	consulInstancer := consul.NewInstancer(
		client, logger, "user_server", Tag, true)

	// 3. 获取endpointer
	// 创建factory 返回的就是 http.NewClient,和直连差不多
	factory := func(urls string) (endpoint.Endpoint, io.Closer, error) {
		target, _ := url.Parse("http://" + urls)
		return transportHttp.NewClient(
			"GET",
			target,
			derectTranspot.EncodeUserRequest,
			derectTranspot.DecodeUserResponse).Endpoint(), nil, nil
	}
	endpointer := sd.NewEndpointer(consulInstancer, factory, logger)

	// 在客服端实现负载均衡，轮询算法
	// mlib := lb.NewRoundRobin(endpointer)
	// 在客户端使用随机选则一个服务端
	mlib := lb.NewRandom(endpointer, time.Now().UnixNano())

	//// 4. 获取全部的endpoints
	//endpoints, _ := endpointer.Endpoints()
	//endP := endpoints[0]
	//fmt.Println("endpoint 长度:", len(endpoints))
	//// 5. 按负载均衡获取其中的一个
	//if len(endpoints) < 0 {
	//	  return
	//}

	for {
		time.Sleep(time.Second * 1)
		endP, _ := mlib.Endpoint()

		// 6.调用服务
		response, err := endP(context.Background(), derectTranspot.UserRequest{ID: 2})
		if err != nil {
			fmt.Println(err)
			return
		}

		// 7.断言响应
		rp := response.(derectTranspot.UserResponse)
		fmt.Printf("recive response : %s\n", rp.Data)
	}
}
