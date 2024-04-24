package logic

import (
	"context"
	"zero/user/code"
	"zero/user/rpc/user"

	"zero/user/api/internal/svc"
	"zero/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.Response, err error) {
	//获取用户的uid
	value := l.ctx.Value("uid").(string)
	if value == "" {
		return nil, code.UserUidNotFoundError
	}
	logout, err := l.svcCtx.User.Logout(l.ctx, &user.LogoutReq{Uid: value})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Code: logout.Code,
		Msg:  logout.Msg,
	}, nil
}
