package redis

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := Client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
