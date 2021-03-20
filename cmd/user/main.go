package main

/*
@Time    : 2021/3/9 14:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/
import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"log"
	rowHttp "net/http"
	"os"
	"os/signal"
	"strconv"
	discovery "sxk.go-kit/api/discovery/user"
	utilty "sxk.go-kit/internal/user"
	ue "sxk.go-kit/internal/user/endpoint"
	us "sxk.go-kit/internal/user/service"
	ut "sxk.go-kit/internal/user/transport"
	"syscall"
)

func main() {
	// use flag receive params
	var id, name, address string
	var port int
	flag.StringVar(&id, "id", "austsxk", "服务唯一ID，如果相同则自动添加唯一区别信息")
	flag.StringVar(&name, "name", "user_server", "服务名称")
	flag.StringVar(&address, "d", "192.168.31.57", "服务唯一ID，如果相同则自动添加唯一区别信息")
	flag.IntVar(&port, "p", 9999, "服务运行的端口")
	flag.Parse()

	discovery.SetServerConfig(id, name, address, port)
	// 1. 初始化实体
	user := us.NewUserImpl()

	// 2. 生成端点
	rateLimit := rate.NewLimiter(1, 3)
	// 使用中间键，添加限流操作
	endpoint1 := utilty.RateLimiterEndpoint(rateLimit)(ue.GenUserEndpoint(user))
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

	log.Printf("user go-kit server is running at %s:%d", address, port)

	// register user server to consul
	errChannel := make(chan error)
	go func() {
		discovery.DiscoveryServer()
		err := rowHttp.ListenAndServe(address+":"+strconv.Itoa(port), router)
		log.Printf("http server is error: %s", err)
		errChannel <- err
	}()

	// signal to kill process
	go func() {
		sigChannel := make(chan os.Signal)
		signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		errChannel <- fmt.Errorf("syscall :%s", <-sigChannel)
	}()

	select {
	case <-errChannel:
		// deregister
		discovery.DeregisterDiscovery()
	}
	fmt.Println("shutdown...")
}
