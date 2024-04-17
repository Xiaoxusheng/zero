package rediscache

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
	ConnMaxIdleTime int
	DialTimeout     int
	ReadTimeout     int
	WriteTimeout    int
	Time            int32
}

// NewRedis  连接redis
func NewRedis(r *Redis) *redis.Client {
	fmt.Println(r)
	rdb := redis.NewClient(&redis.Options{
		Addr:            r.Addr,
		Password:        r.Password, // 没有密码，默认值
		DB:              r.DB,
		PoolSize:        r.PoolSize,
		MinIdleConns:    r.MinIdleConns,
		MaxIdleConns:    r.MaxIdleConns,
		ReadTimeout:     time.Duration(r.ReadTimeout) * time.Second,
		WriteTimeout:    time.Duration(r.WriteTimeout) * time.Second,
		ConnMaxIdleTime: time.Duration(r.ConnMaxIdleTime) * time.Second,
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
