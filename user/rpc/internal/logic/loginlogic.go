package logic

import (
	"context"
	"fmt"

	"zero/user/rpc/internal/svc"
	"zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRes) (*user.Response, error) {
	// todo: add your logic here and delete this line
	fmt.Println(in.Username, in.Password)
	return &user.Response{}, nil
}
