// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package users

import (
	"context"

	"zero/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRes     = user.LoginRes
	LoginResp    = user.LoginResp
	LogoutReq    = user.LogoutReq
	RegisterReq  = user.RegisterReq
	RegisterResp = user.RegisterResp
	Response     = user.Response
	UserInfoReq  = user.UserInfoReq
	UserInfoResp = user.UserInfoResp

	Users interface {
		// 用户登录
		Login(ctx context.Context, in *LoginRes, opts ...grpc.CallOption) (*LoginResp, error)
		// 用户注册
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		// 用户信息
		GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
		// 退出登录
		Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUsers struct {
		cli zrpc.Client
	}
)

func NewUsers(cli zrpc.Client) Users {
	return &defaultUsers{
		cli: cli,
	}
}

// 用户登录
func (m *defaultUsers) Login(ctx context.Context, in *LoginRes, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUsersClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 用户注册
func (m *defaultUsers) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUsersClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 用户信息
func (m *defaultUsers) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := user.NewUsersClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

// 退出登录
func (m *defaultUsers) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUsersClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}
