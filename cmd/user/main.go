package main

/*
@Time    : 2021/3/9 14:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/
import (
	"fmt"
	"github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"log"
	rowHttp "net/http"
	"os"
	"os/signal"
	discovery "sxk.go-kit/api/discovery/user"
	ue "sxk.go-kit/internal/user/endpoint"
	us "sxk.go-kit/internal/user/service"
	ut "sxk.go-kit/internal/user/transport"
	"syscall"
)

func main() {
	// 1. 初始化实体
	user := us.NewUserImpl()

	// 2. 生成端点
	endpoint1 := ue.GenUserEndpoint(user)
	endpoint2 := ue.GenAddUserEndpoint(user)

	// 3. 使用kit内置的http定义服务
	server1 := http.NewServer(endpoint1, ut.EncodeUserRequest, ut.EncodeUserResponse)
	server2 := http.NewServer(endpoint2, ut.EncodeAddUserRequest, ut.EncodeAddUserResponse)

	// 4. 使用第三方工具包装路由
	router := mymux.NewRouter()
	router.Methods("Get").Path("/user/{id:\\d+}").Handler(server1)

	// 5. add user methods
	router.Methods("Post").Path("/user/add").Handler(server2)

	// 6. add a health check
	router.Methods("Get").Path("/health").HandlerFunc(func(writer rowHttp.ResponseWriter, request *rowHttp.Request) {
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status": "ok"}`))
	})

	log.Println("user go-kit server is running at 127.0.0.1:9999")

	// register user server to consul
	errChannel := make(chan error)
	go func() {
		discovery.DiscoveryServer()
		err := rowHttp.ListenAndServe(":9999", router)
		log.Printf("http server is error: %s", err)
		errChannel <- err
	}()

	// signal to kill process
	go func() {
		sigChannel := make(chan os.Signal)
		signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
		errChannel <- fmt.Errorf("syscall :%s", <-sigChannel)
	}()

	<-errChannel
	// deregister
	discovery.DeregisterDiscovery()
}
