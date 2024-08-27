package ioc

import (
	"context"
	"github.com/Teresajw/go_project/webook/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {
	rd := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Pwd,
		DB:       config.Config.Redis.DB,
	})
	err := rd.Ping(context.Background()).Err()
	if err != nil {
		panic("failed to connect redis")
	}
	return rd
}
