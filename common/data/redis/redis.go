package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
)

//func init() {
//	// 初始化redis
//	initRedis()
//}

var RedisClient *redis.Client

// 初始化redis连接
func InitRedis(redisConfig *RedisConfig) {
	log.Println("连接redis", redisConfig.Addr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
	})
	RedisClient = rdb
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Panicln("Redis client ping failed: ", err)
	}
	info := RedisClient.Info(context.Background())
	// 获取redis版本号
	versionIndex := strings.Index(info.String(), "redis_version:")
	version := info.String()[versionIndex+14 : versionIndex+20]
	log.Println("Redis client initialized successfully. Redis version: ", version)

}

type RedisConfig struct {
	Addr     string
	Password string
}
