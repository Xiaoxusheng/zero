package logic

import (
	"context"
	"encoding/base64"
	"net/http"
	"zero/pkg/rediscache"
	"zero/user/code"
	"zero/user/model"
	"zero/utils"

	"zero/user/rpc/internal/svc"
	"zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRes) (*user.RegisterReq, error) {
	//检测电话号码是否已经存在
	u, err := l.svcCtx.UserModel.FindPhone(l.ctx, in.Phone)
	if err != nil || u == nil {
		l.Logger.Error(err.Error())
		return nil, code.UserExistError
	}
	//检测名称
	u1, err := l.svcCtx.UserModel.FindName(l.ctx, in.Name)
	if err != nil || u1 == nil {
		l.Logger.Error(err.Error())
		return nil, code.UserNameExistError
	}
	//生成随机盐值
	salt, err := utils.Salt(16)
	if err != nil {
		l.Logger.Error(err.Error())
		return nil, code.UserNameExistError
	}
	//加盐加密
	pwd := utils.HashPassword(in.Password, salt)
	uid := utils.GetUidV4()
	//插入数据库
	err = l.svcCtx.UserModel.InsertUser(l.ctx, &model.User{
		Name:            in.Name,
		Password:        pwd,
		Phone:           in.Phone,
		Slat:            base64.URLEncoding.EncodeToString(salt),
		Sex:             in.Sex,
		Avatar:          in.Avatar,
		BackgroundImage: in.BackgroundImage,
		Uid:             uid,
	})
	if err != nil {
		return nil, err
	}
	//写入缓存
	l.svcCtx.Redis.HSet(l.ctx, uid, rediscache.Salt, base64.URLEncoding.EncodeToString(salt))

	return &user.RegisterReq{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
