package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api/echo"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
	"github.com/cloudwego/kitex/client"
)

func main() {
	//创建客户端，指定调用地址
	cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("127.0.0.1:8888"),
		//传递客户端信息
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}),
	)
	//如果创建错误，直接panic
	if err != nil {
		panic(err)
	}
	//客户端使用context发送带有hello消息的请求
	res, err := cli.Echo(context.TODO(), &api.Request{Message: "hello"})
	//如果请求错误，打印错误
	if err != nil {
		fmt.Println(err)
	}
	//打印请求结果
	fmt.Printf("res: %v\n", res)
}
