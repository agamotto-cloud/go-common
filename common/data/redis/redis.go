package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"strings"
)

//func init() {
//	// 初始化redis
//	initRedis()
//}

var RedisClient *redis.Client

// 初始化redis连接
func InitRedis(redisConfig *RedisConfig) {
	log.Info().Msgf("连接redis %s", redisConfig.Addr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
	})
	RedisClient = rdb
	go func() {
		_, err := RedisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Panic().Err(err).Msg("Redis client ping failed")
		}
		info := RedisClient.Info(context.Background())
		// 获取redis版本号
		versionIndex := strings.Index(info.String(), "redis_version:")
		version := info.String()[versionIndex+14 : versionIndex+20]
		log.Info().Msgf("Redis client initialized successfully. Redis version: %s", version)
	}()
}

type RedisConfig struct {
	Addr     string
	Password string
}
