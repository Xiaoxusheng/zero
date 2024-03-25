package logic

import (
	"context"
	"fmt"
	"zero/user/rpc/user"

	"zero/user/api/internal/svc"
	"zero/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRes) (resp *types.Response, err error) {
	fmt.Println(req)
	// todo: add your logic here and delete this line
	login, err := l.svcCtx.User.Login(l.ctx, &user.LoginRes{
		Username: "1",
		Password: "2",
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Code: login.Code,
		Msg:  login.Msg,
		Data: nil,
	}, nil
}
