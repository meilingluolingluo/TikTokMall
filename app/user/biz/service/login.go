package service

import (
	"context"
	"errors"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
}

func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// 参数校验
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	// 获取用户信息
	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}
	// 密码比对
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	// 组织响应
	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}
	return resp, nil
}
