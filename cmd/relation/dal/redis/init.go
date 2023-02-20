package redis

import (
	"douyin/pkg/constants"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// Init init DB
func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddress,
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	RedisClient = rdb
}
