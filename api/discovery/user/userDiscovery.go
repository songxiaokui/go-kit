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
	"fmt"
	"github.com/google/uuid"
	mapi "github.com/hashicorp/consul/api"
	"log"
	"strconv"
	"sync"
)

const (
	DefaultAddress = "192.168.31.57"
)

var (
	UClient        *mapi.Client
	once           sync.Once
	serviceID      string
	serviceName    string
	serviceAddress string
	ServicePort    int
)

func init() {
	once.Do(func() {
		// consul address
		config := mapi.DefaultConfig()
		config.Address = DefaultAddress + ":8500"
		client, err := mapi.NewClient(config)
		if err != nil {
			log.Fatal(err)
		}
		UClient = client
	})
}

func SetServerConfig(id, name, address string, port int) {
	serviceID = id + uuid.New().String()
	serviceName = name
	serviceAddress = address
	ServicePort = port
}

func DiscoveryServer() {
	// server struct
	register := mapi.AgentServiceRegistration{}
	register.ID = serviceID
	register.Name = serviceName
	register.Address = serviceAddress
	register.Port = ServicePort
	register.Tags = []string{"userModel"}

	// server check struct
	check := mapi.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = fmt.Sprintf("http://%s:%s/health", serviceAddress, strconv.Itoa(ServicePort))

	register.Check = &check

	err := UClient.Agent().ServiceRegister(&register)
	if err != nil {
		log.Fatal(err)
	}

}

// Deregister
func DeregisterDiscovery() {
	err := UClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		log.Fatal(err)
	}
}
