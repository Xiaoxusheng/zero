package svc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"zero/pkg/orm"
	"zero/pkg/rediscache"
	"zero/user/model"
	"zero/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserMode
	Redis     *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println(c.DB)
	db, err := orm.NewMysql(&orm.Config{
		Dsn:             c.DB.Dsn,
		MaxIdleConns:    c.DB.MaxIdleConns,
		MaxOpenConns:    c.DB.MaxOpenConns,
		ConnMaxLifetime: c.DB.ConnMaxLifetime,
	})
	if err != nil {
		panic(err)
	}

	r := rediscache.NewRedis(&rediscache.Redis{
		Addr:            c.RedisCache.Addr,
		Password:        c.RedisCache.Password,
		DB:              c.RedisCache.DB,
		PoolSize:        c.RedisCache.PoolSize,
		MinIdleConns:    c.RedisCache.MinIdleConns,
		MaxIdleConns:    c.RedisCache.MaxIdleConns,
		ConnMaxIdleTime: c.RedisCache.ConnMaxIdleTime,
		DialTimeout:     c.RedisCache.DialTimeout,
		ReadTimeout:     c.RedisCache.ReadTimeout,
		WriteTimeout:    c.RedisCache.WriteTimeout,
		Time:            c.RedisCache.Time,
	})

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUser(db.DB),
		Redis:     r,
	}
}
