package logic

import (
	"context"
	"zero/user/rpc/user"
	"zero/utils"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	login, err := l.svcCtx.User.Login(l.ctx, &user.LoginRes{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	token := utils.GetToken(login.Uid, l.svcCtx.Config.Auth.AccessSecret, int32(l.svcCtx.Config.Auth.AccessExpire))
	return &types.LoginResp{
		Token:           token,
		Name:            login.Name,
		Sex:             login.Sex,
		Phone:           login.Phone,
		BackgroundImage: login.BackgroundImage,
		Avatar:          login.Avatar,
		Code:            login.Code,
		Msg:             login.Msg,
	}, nil
}
