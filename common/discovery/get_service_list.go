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
func GetServiceList[T any](ctx context.Context, serviceName string) []ServerNode[T] {
	serverKey := "service:" + serviceName
	result, err := redis.RedisClient.HGetAll(ctx, serverKey).Result()
	if err != nil {
		log.Error().Msgf("获取服务列表失败 %s", err.Error())
		return []ServerNode[T]{}
	}
	serverList := mapsToServerList[T](result)
	return serverList
}

// GetServiceInfo 获取指定服务的具体节点信息
func GetServiceInfo[T any](ctx context.Context, serviceName string, address string, port int) ServerNode[T] {
	serverKey := "service:" + serviceName
	result, err := redis.RedisClient.HGet(ctx, serverKey, address+":"+strconv.Itoa(port)).Result()
	if err != nil {
		log.Error().Msgf("获取服务列表失败 %s", err.Error())
		return ServerNode[T]{}
	}
	var serverNode ServerNode[T]
	err = json.Unmarshal([]byte(result), &serverNode)
	if err != nil {
		log.Error().Msgf("json转换失败 %s", err.Error())
		return ServerNode[T]{}
	}
	return serverNode
}

// GetLocalServiceList 获取本服务的服务列表
func GetLocalServiceList[T any](ctx context.Context) []ServerNode[T] {
	serverKey := "service:" + config.GetServerConfig().Name
	result, err := redis.RedisClient.HGetAll(ctx, serverKey).Result()
	if err != nil {
		log.Error().Msgf("获取服务列表失败 %s", err.Error())
		return []ServerNode[T]{}
	}
	serverList := mapsToServerList[T](result)
	return serverList
}

// GetLocalServiceInfo 获取本服务的服务信息
func GetLocalServiceInfo[T any]() ServerNode[T] {
	//serverNodeInfo.Info需要转成T类型
	var serInfo ServerNode[T]
	serInfo = ServerNode[T]{
		Info: serverNodeInfo.Info.(T),
		//剩下的属性全过来
		Address:        serverNodeInfo.Address,
		Port:           serverNodeInfo.Port,
		ActiveLastTime: serverNodeInfo.ActiveLastTime,
	}
	return serInfo

}
