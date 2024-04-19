package logic

import (
	"context"
	"zero/user/rpc/user"

	"zero/user/api/internal/svc"
	"zero/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.Response, err error) {
	//注册
	r, err := l.svcCtx.User.Register(l.ctx, &user.RegisterRes{
		Name:            req.Username,
		Password:        req.Password,
		Phone:           req.Phone,
		Sex:             req.Sex,
		BackgroundImage: req.BackgroundImage,
		Avatar:          req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Code: r.Code,
		Msg:  r.Msg,
	}, nil
}
