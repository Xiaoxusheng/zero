package logic

import (
	"context"
	"net/http"
	"zero/user/code"

	"zero/user/rpc/internal/svc"
	"zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Logout 退出登录
func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.Response, error) {
	_, err := l.svcCtx.Redis.Del(l.ctx, in.Uid).Result()
	if err != nil {
		return nil, err
	}
	return &user.Response{
		Code: http.StatusOK,
		Msg:  code.SuccessCode,
	}, nil
}
