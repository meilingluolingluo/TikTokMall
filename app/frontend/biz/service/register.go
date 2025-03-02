package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"

	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", "1")
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
