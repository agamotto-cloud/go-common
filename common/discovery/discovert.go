package discovery

import (
	"context"
	"encoding/json"
	"github.com/agamotto-cloud/go-common/common/config"
	"github.com/agamotto-cloud/go-common/common/data/redis"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type ServerNode struct {
	ActiveLastTime int64       `json:"activeLastTime"`
	Address        string      `json:"Address"`
	Port           int         `json:"port"`
	Info           interface{} `json:"info"`
}

var serverNodeInfo = ServerNode{}

func init() {
	serverNodeInfo = ServerNode{
		ActiveLastTime: time.Now().Unix(),
		Address:        getLocalIP(),
		Port:           8080,
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-time.Tick(5 * time.Second):
				// 每5s执行一次任务
				updateServerNode()
				break

			case <-sigs:
				// 接收到关闭信号
				closeServer()
				log.Println("停止服务")
				go func() {
					os.Exit(0)
				}()
				return
			}
		}

	}()

}

// 一个定时任务，每5s执行一次，
func updateServerNode() {
	serverConfig := config.GetServerConfig()
	serverNodeInfo.ActiveLastTime = time.Now().Unix()
	serverNodeInfo.Port = serverConfig.Port
	if getServerInfoFunc != nil {
		serverNodeInfo.Info = getServerInfoFunc(serverNodeInfo)
	}
	jsonStr, _ := json.Marshal(serverNodeInfo)
	log.Println("service discover info :", string(jsonStr))
	serverKey := "service:" + serverConfig.Name
	redis.RedisClient.HSet(context.Background(), serverKey, serverNodeInfo.Address+":"+strconv.Itoa(serverNodeInfo.Port), jsonStr)
	redis.RedisClient.Expire(context.Background(), serverKey, time.Hour)
	result, err := redis.RedisClient.HGetAll(context.Background(), serverKey).Result()
	if err != nil {
		log.Println("获取服务列表失败", err.Error())
		return
	}
	serverList := mapsToServerList(result)
	//剔除掉超时的服务 超时时间为10分钟
	for i, v := range serverList {
		if time.Now().Unix()-v.ActiveLastTime > 600 {
			redis.RedisClient.HDel(context.Background(), serverKey, v.Address+":"+strconv.Itoa(v.Port))
			serverList = append(serverList[:i], serverList[i+1:]...)
		}
	}
}

// 服务关闭的处理
func closeServer() {
	serverKey := "service:" + config.GetServerConfig().Name
	redis.RedisClient.HDel(context.Background(), serverKey, serverNodeInfo.Address+":"+strconv.Itoa(serverNodeInfo.Port))
}

// 将map转换为ServerNode数组
func mapsToServerList(maps map[string]string) []ServerNode {
	var serverList = make([]ServerNode, 0)
	for _, v := range maps {
		var serverNode ServerNode
		err := json.Unmarshal([]byte(v), &serverNode)
		if err != nil {
			log.Println("json转换失败", err.Error())
			continue
		}
		serverList = append(serverList, serverNode)
	}
	return serverList
}

var getServerInfoFunc func(serverNode ServerNode) interface{}

// SetServerInfoFunc 设置服务信息的函数，参数是一个函数,保存这个函数,在updateServerNode方法中会定时调用这个函数
func SetServerInfoFunc(f func(serverNode ServerNode) interface{}) {
	getServerInfoFunc = f
}

// 获取本机IP
func getLocalIP() string {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range adders {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}
