package data

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
)

func init() {
	// 初始化redis
	initRedis()
}

var RedisClient *redis.Client

// 初始化redis连接
func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-17517.c299.asia-northeast1-1.gce.cloud.redislabs.com:17517",
		Password: "DongAgamotto!@#",
	})
	RedisClient = rdb
	info := RedisClient.Info(context.Background())
	// 获取redis版本号
	versionIndex := strings.Index(info.String(), "redis_version:")
	version := info.String()[versionIndex+14 : versionIndex+20]
	log.Println("Redis client initialized successfully. Redis version: ", version)
}
