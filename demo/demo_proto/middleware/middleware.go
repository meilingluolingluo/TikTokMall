package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func MiddleWare(next endpoint.Endpoint) endpoint.Endpoint {
	//记录每个请求的执行时间
	return func(ctx context.Context, req, resp interface{}) (err error) {
		begin := time.Now()
		err = next(ctx, req, resp)
		fmt.Println(time.Since(begin))
		return err

	}
}
