package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"

	"github.com/bytedance/gopkg/cloud/metainfo"

	pbapi "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	// Finish your business logic.
	//从上下文获取CLIENT_NAME元信息
	clientNAME, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME")
	fmt.Println(clientNAME, ok)
	//返回业务异常

	if req.Message == "error" {
		//返回100类，4001信息的异常码
		return nil, kerrors.NewGRPCBizStatusError(1004001, "client param error")
	}
	return &pbapi.Response{
		Message: req.Message,
	}, nil
}
