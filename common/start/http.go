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
	cc, ss := logger.CreateSpan(context.Background(), "http server")
	defer ss.End()
	log := logger.GetLog(cc)
	serverConfig := config.GetServerConfig()
	log.Info().Int("port", serverConfig.Port).Msg("server start")
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(Logger())
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debug().Msgf("endpoint %v %v %v ", httpMethod, absolutePath, handlerName)
	}
	_ = router.SetTrustedProxies(nil)
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(serverConfig.Port),
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("listen")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server Shutdown")
	}
	log.Info().Msg("Server exiting")
}

// 记录每一个请求的日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//t := time.Now()
		ctt, ss := logger.CreateSpan(c, c.Request.URL.Path)
		defer ss.End()
		// 请求前
		log := logger.GetLog(ctt)
		log.Info().Msgf("请求接口 %s", c.Request.URL.Path)
		c.Next()

		//// 请求后
		//latency := time.Since(t)
		//status := c.Writer.Status()
		//log.Info("请求接口",
	}
}
