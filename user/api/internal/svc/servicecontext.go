package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero/user/api/internal/config"
	"zero/user/rpc/users"
)

type ServiceContext struct {
	Config config.Config
	User   users.Users
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   users.NewUsers(zrpc.MustNewClient(c.User)),
	}
}
