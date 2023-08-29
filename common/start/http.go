package start

import (
	"context"
	"github.com/agamotto-cloud/go-common/common/config"
	"github.com/agamotto-cloud/go-common/common/logger"
	"github.com/agamotto-cloud/go-common/common/response"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// HttpServer http服务 参数是一个回调，可以在回调中注册路由
func HttpServer(routerReg func(r *gin.Engine)) {

	serverConfig := config.GetServerConfig()
	log.Info().Int("port", serverConfig.Port).Msg("server start")
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(loggerHandle())
	router.NoRoute(func(c *gin.Context) {
		response.Error(response.NotFoundError, c)
	})
	if routerReg != nil {
		routerReg(router)
	}
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

// Logger 记录每一个请求的日志
func loggerHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		ctt, ss := logger.CreateSpan(c.Request.Context(), c.Request.URL.Path)
		c.Request = c.Request.WithContext(ctt)
		defer ss.End()
		// 请求前
		log1 := log.Ctx(c.Request.Context())
		log1.Info().Msgf("请求接口 %s", c.Request.URL.Path)
		c.Next()

		//// 请求后
		latency := time.Since(t)
		status := c.Writer.Status()
		log1.Info().Int("status", status).
			Dur("latency", latency).
			Msgf("请求接口 %s 结果 %d 耗时 %s", c.Request.URL.Path, status, latency.String())
	}
}
