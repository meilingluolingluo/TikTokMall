package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/rpcinfo"

	api "github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *api.Request) (resp *api.Response, err error) {
	// Finish your business logic.
	//获取rpc信息
	info := rpcinfo.GetRPCInfo(s.ctx)
	fmt.Println(info.From().ServiceName())
	//返回响应对象resp
	return &api.Response{
		Message: req.Message,
	}, nil
}
