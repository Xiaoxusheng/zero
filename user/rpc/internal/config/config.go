package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		Dsn             string
		MaxIdleConns    int   `json:",default=100"`
		MaxOpenConns    int   `json:",default=100"`
		ConnMaxLifetime int64 `json:",default=3600"`
	}
}
