package config

import (
	"context"
	"flag"
	"github.com/agamotto-cloud/go-common/common/data"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// 这个是全部配置
var ConfigProps map[string]interface{}

// 这个是格式化后的配置
var ConfigMap = make(map[string]interface{})

// 首先从配置文件中获取配置，然后根据配置文件去建立redis连接，然后再去获取redis中的配置，最后将两个配置合并
func init() {
	name := flag.String("env", "dev", "a string")
	flag.Parse()
	env := *name
	if env == "" {
		env = "dev"
	}
	log.Println("加载服务器配置 ", env)

	// 读取本地配置文件
	configData, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}

	err = yaml.Unmarshal(configData, &ConfigProps)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	//初始化redis
	redisConfig := GetConfig("redis", data.RedisConfig{})
	data.InitRedis(redisConfig)

	//从命令行参数中获取当前运行环境

	configKey := "config:service:admin-server:" + env
	// 从redis中获取配置
	redisConfigData, err := data.RedisClient.Get(context.Background(), configKey).Result()
	if err != nil {
		log.Println("读取redis配置文件失败: ", configKey, err)
		return
	}

	var redisMap map[string]interface{}
	err = yaml.Unmarshal([]byte(redisConfigData), &redisMap)
	if err != nil {
		log.Fatalf("读取redis配置文件失败: %v", err)
	}
	// 将redis中的配置和本地配置合并
	for k, v := range redisMap {
		ConfigProps[k] = v
	}

}

func GetServerConfig() *ServerConfig {
	return GetConfig("server", ServerConfig{})
}

func GetConfig[T any](configKey string, serverConfig T) *T {
	//如果之前获取过配置，直接返回
	if ConfigMap[configKey] != nil {
		configData := ConfigMap[configKey].(T)
		return &configData
	}

	err := mapstructure.Decode(ConfigProps[configKey], &serverConfig)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	ConfigMap[configKey] = serverConfig
	return &serverConfig
}

type ServerConfig struct {
	Port int
}
