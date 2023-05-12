package start

import (
	"context"
	"github.com/agamotto-cloud/go-common/common/config"
	"github.com/agamotto-cloud/go-common/common/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func HttpServer() {
	log := logger.GetLog(nil)
	serverConfig := config.GetServerConfig()
	log.Info("server start")
	router := gin.New()
	router.Use(gin.Recovery())
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debug("endpoint %v %v %v \n", httpMethod, absolutePath, handlerName)
	}
	_ = router.SetTrustedProxies(nil)
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(serverConfig.Port),
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}
