package discovery

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"server/pkg/data"
	"strconv"
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

	go func() {
		for {
			select {
			case <-time.Tick(5 * time.Second):
				// 每5s执行一次任务
				updateServerNode()
			}
		}
	}()

}

// 一个定时任务，每5s执行一次，
func updateServerNode() {
	serverNodeInfo.ActiveLastTime = time.Now().Unix()
	if getServerInfoFunc != nil {
		serverNodeInfo.Info = getServerInfoFunc(serverNodeInfo)
	}
	jsonStr, _ := json.Marshal(serverNodeInfo)
	log.Println("service discover info :", string(jsonStr))
	data.RedisClient.HSet(context.Background(), "service:admin", serverNodeInfo.Address+":"+strconv.Itoa(serverNodeInfo.Port), jsonStr)
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
