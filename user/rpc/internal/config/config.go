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
	RedisCache struct {
		Addr            string
		Password        string
		DB              int
		PoolSize        int
		MinIdleConns    int
		MaxIdleConns    int
		ConnMaxIdleTime int
		DialTimeout     int
		ReadTimeout     int
		WriteTimeout    int
		Time            int32
	}
}
