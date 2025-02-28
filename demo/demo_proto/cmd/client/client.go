package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/middleware"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/kitex/pkg/kerrors"

	"github.com/cloudwego/kitex/pkg/transmeta"

	"github.com/cloudwego/kitex/transport"

	"github.com/bytedance/gopkg/cloud/metainfo"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	_ "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	// 创建 Consul 注册中心
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}
	// 创建客户端
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.MiddleWare),
	)

	if err != nil {
		log.Fatal(err)
	}
	//构造带元信息的上下文
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	// 使用客户端发送请求
	res, err := c.Echo(ctx, &pbapi.Request{Message: "Hello"})

	var bizErr *kerrors.GRPCBizStatusError
	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Printf("%#v", bizErr)
		}
		klog.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %v\n", res)
}
