package cache

import "C"
import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	Addr            string
	Password        string
	DB              int
	PoolSize        int
	MinIdleConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	Time            int32
}

// NewRedis  连接redis
func NewRedis(r *Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:            r.Addr,
		Password:        r.Password, // 没有密码，默认值
		DB:              r.DB,
		PoolSize:        r.PoolSize,
		MinIdleConns:    r.MinIdleConns,
		MaxIdleConns:    r.MaxIdleConns,
		ReadTimeout:     r.ReadTimeout,
		WriteTimeout:    r.WriteTimeout,
		ConnMaxIdleTime: r.ConnMaxIdleTime * time.Second,
	}).WithTimeout(time.Duration(r.Time) * time.Second)
	ctx := context.Background()
	res, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if res == "PONG" {
		fmt.Println("redis连接成功!")
	}
	return rdb
}
