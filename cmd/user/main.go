package main

/*
@Time    : 2021/3/9 14:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main.go
@Software: GoLand
*/
import (
	"github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"log"
	rowHttp "net/http"
	ue "sxk.go-kit/internal/user/endpoint"
	us "sxk.go-kit/internal/user/service"
	ut "sxk.go-kit/internal/user/transport"
)

func main() {
	// 1. 初始化实体
	user := us.NewUserImpl()

	// 2. 生成端点
	endpoint := ue.GenUserEndpoint(user)

	// 3. 使用kit内置的http定义服务
	server := http.NewServer(endpoint, ut.EncodeUserRequest, ut.EncodeUserResponse)

	// 4. 使用第三方工具包装路由
	router := mymux.NewRouter()
	router.Methods("Get").Path("/user/{id:\\d+}").Handler(server)
	log.Println("user go-kit server is running at 127.0.0.1:9999")

	err := rowHttp.ListenAndServe(":9999", router)
	if err != nil {
		log.Fatal("servers error:", err)
	}

}
