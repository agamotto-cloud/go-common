package main

import (
	"context"
	"encoding/json"
	data "github.com/agamotto-cloud/go-common/common/data/db"
	"github.com/agamotto-cloud/go-common/common/discovery"
	_ "github.com/agamotto-cloud/go-common/common/discovery"
	"github.com/agamotto-cloud/go-common/common/start"
)

func main() {
	// 加载配置文件
	//data.Init()
	r := discovery.GetServiceList(context.Background(), "admin")
	for _, node := range r {
		//用json的格式打印node
		d, _ := json.Marshal(node)
		println(string(d))
	}

	start.HttpServer(nil)

	data.GlobalDB.Exec("select 1 from dual")
	//logger.GetLogger().Info("server exit")
}
