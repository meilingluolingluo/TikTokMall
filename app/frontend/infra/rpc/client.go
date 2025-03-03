package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"

	consul "github.com/kitex-contrib/registry-consul"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}
