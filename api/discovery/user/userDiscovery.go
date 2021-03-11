package user

/*
@Time    : 2021/3/11 14:03
@Author  : austsxk
@Email   : austsxk@163.com
@File    : userDiscovery.go
@Software: GoLand
*/

// automatic discovery user server

import (
	mapi "github.com/hashicorp/consul/api"
	"log"
)

func DiscoveryServer() {
	// consul address
	config := mapi.DefaultConfig()
	config.Address = "192.168.31.102:8500"

	// server struct
	register := mapi.AgentServiceRegistration{}
	register.ID = "austsxk.user.v1"
	register.Name = "user_server"
	register.Address = "192.168.31.102"
	register.Port = 9999
	register.Tags = []string{"userModel"}

	// server check struct
	check := mapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://192.168.31.102:9999/health"

	register.Check = &check

	client, err := mapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Agent().ServiceRegister(&register)
	if err != nil {
		log.Fatal(err)
	}

}
