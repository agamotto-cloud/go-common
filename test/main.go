package main

import (
	data "github.com/agamotto-cloud/go-common/common/data/db"
	"github.com/agamotto-cloud/go-common/common/start"
)

func main() {
	// 加载配置文件
	//data.Init()

	start.HttpServer()

	data.GlobalDB.Exec("select 1 from dual")
	//logger.GetLogger().Info("server exit")
}
