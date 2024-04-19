package logic

import (
	"context"
	"encoding/base64"
	"net/http"
	"zero/pkg/rediscache"
	"zero/user/code"
	"zero/utils"

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

// Login 用户登录
func (l *LoginLogic) Login(in *user.LoginRes) (*user.LoginResp, error) {
	//判断用户是否存在
	userinfo, err := l.svcCtx.UserModel.FindName(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	//判断密码是否正确
	salt := l.svcCtx.Redis.HGet(l.ctx, userinfo.Uid, rediscache.Salt).Val()
	decodeString, err := base64.URLEncoding.DecodeString(salt)
	if err != nil {
		return nil, err
	}
	pwd := utils.HashPassword(in.Password, decodeString)
	if pwd != userinfo.Password {
		return nil, code.UserPwdError
	}
	//获取token
	return &user.LoginResp{
		Uid:             userinfo.Uid,
		Name:            userinfo.Name,
		Phone:           userinfo.Phone,
		Sex:             userinfo.Sex,
		BackgroundImage: userinfo.BackgroundImage,
		Avatar:          userinfo.Avatar,
		Code:            http.StatusOK,
		Msg:             code.SuccessCode,
	}, nil
}
