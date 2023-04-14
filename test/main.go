package main

import (
	"github.com/agamotto-cloud/go-common/common/config"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	// 加载配置文件
	//data.Init()

	serverConfig := config.GetServerConfig()
	log.Println("server start")
	r := gin.Default()
	err := r.Run(":" + strconv.Itoa(serverConfig.Port))
	if err != nil {
		log.Fatal("server start error", err.Error())
	}
}
