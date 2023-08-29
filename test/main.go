package main

import (
	"buf.build/gen/go/agamotto/test/protocolbuffers/go/proto/user"
	"context"
	"encoding/json"
	data "github.com/agamotto-cloud/go-common/common/data/db"
	"github.com/agamotto-cloud/go-common/common/discovery"
	_ "github.com/agamotto-cloud/go-common/common/discovery"
	"github.com/agamotto-cloud/go-common/common/start"
	"github.com/agamotto-cloud/go-common/test/rpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	// 加载配置文件
	//data.Init()
	discovery.SetServerInfoFunc(func(node discovery.ServerNode[any]) interface{} {
		return map[string]any{
			"name": "test",
		}
	})
	userRequest := &user.CreateUserRequest{
		Username: "test",
	}
	log.Info().Msgf("userRequest %v", userRequest)
	//api := discovery.GetServiceList[any](context.Background(), "test")
	//CreateUserRequest := struct {
	r := discovery.GetServiceList[any](context.Background(), "test")
	for _, node := range r {
		//用json的格式打印node
		d, _ := json.Marshal(node)
		println(string(d))
	}
	go start.RpcServer(func(s *grpc.Server) {
	})
	go rpc.CallRpc()
	start.HttpServer(nil)

	data.GlobalDB.Exec("select 1 from dual")

	//logger.GetLogger().Info("server exit")
}
