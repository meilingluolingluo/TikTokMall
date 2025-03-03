package service

import (
	"context"

	"github.com/joho/godotenv"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"

	"testing"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return
	}
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	req := &user.RegisterReq{
		Email:           "test@test.com",
		Password:        "123456",
		ConfirmPassword: "12345",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
