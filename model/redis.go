package model

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatal("Redis连接失败: ", err)
	}
	log.Println("Redis已连接")
}
