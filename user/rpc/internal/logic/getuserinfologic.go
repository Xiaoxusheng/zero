package logic

import (
	"context"

	"zero/user/rpc/internal/svc"
	"zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserInfo 用户信息
func (l *GetUserInfoLogic) GetUserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoResp{}, nil
}
