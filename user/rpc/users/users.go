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
	LoginRes    = user.LoginRes
	RegisterReq = user.RegisterReq
	RegisterRes = user.RegisterRes
	Response    = user.Response

	Users interface {
		// 用户登录
		Login(ctx context.Context, in *LoginRes, opts ...grpc.CallOption) (*Response, error)
		// 用户注册
		Register(ctx context.Context, in *RegisterRes, opts ...grpc.CallOption) (*RegisterReq, error)
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
func (m *defaultUsers) Login(ctx context.Context, in *LoginRes, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUsersClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 用户注册
func (m *defaultUsers) Register(ctx context.Context, in *RegisterRes, opts ...grpc.CallOption) (*RegisterReq, error) {
	client := user.NewUsersClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}
