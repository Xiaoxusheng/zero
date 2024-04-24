package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero/video/api/internal/config"
	"zero/video/rpc/videoserver"
)

type ServiceContext struct {
	Config config.Config
	Video  videoserver.VideoServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Video:  videoserver.NewVideoServer(zrpc.MustNewClient(c.Video)),
	}
}
