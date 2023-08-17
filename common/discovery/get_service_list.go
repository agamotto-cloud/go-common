package discovery

import (
	"context"
	"encoding/json"
	"github.com/agamotto-cloud/go-common/common/config"
	"github.com/agamotto-cloud/go-common/common/data/redis"
	"github.com/rs/zerolog/log"
	"strconv"
)

// GetServiceList 获取指定服务的节点列表
func GetServiceList(ctx context.Context, serviceName string) []ServerNode {
	serverKey := "service:" + serviceName
	result, err := redis.RedisClient.HGetAll(ctx, serverKey).Result()
	if err != nil {
		log.Error().Msgf("获取服务列表失败 %s", err.Error())
		return []ServerNode{}
	}
	serverList := mapsToServerList(result)
	return serverList
}

// GetServiceInfo 获取指定服务的具体节点信息
func GetServiceInfo(ctx context.Context, serviceName string, address string, port int) ServerNode {
	serverKey := "service:" + serviceName
	result, err := redis.RedisClient.HGet(ctx, serverKey, address+":"+strconv.Itoa(port)).Result()
	if err != nil {
		log.Error().Msgf("获取服务列表失败 %s", err.Error())
		return ServerNode{}
	}
	var serverNode ServerNode
	err = json.Unmarshal([]byte(result), &serverNode)
	if err != nil {
		log.Error().Msgf("json转换失败 %s", err.Error())
		return ServerNode{}
	}
	return serverNode
}

// GetLocalServiceList 获取本服务的服务列表
func GetLocalServiceList(ctx context.Context) []ServerNode {
	serverKey := "service:" + config.GetServerConfig().Name
	result, err := redis.RedisClient.HGetAll(ctx, serverKey).Result()
	if err != nil {
		log.Error().Msgf("获取服务列表失败 %s", err.Error())
		return []ServerNode{}
	}
	serverList := mapsToServerList(result)
	return serverList
}

// GetLocalServiceInfo 获取本服务的服务信息
func GetLocalServiceInfo() ServerNode {
	return serverNodeInfo
}
