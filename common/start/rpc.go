package start

import (
	"context"
	"github.com/agamotto-cloud/go-common/common/logger"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"time"
)

// HttpServer 启动http服务
func RpcServer(srvReg func(srv *grpc.Server)) {
	//设置监听9000端口,添加统一前置处理器
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(unaryInterceptor),
	}

	srv := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Info().Msg("server start on 9000")
	err = srv.Serve(lis)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 生成请求的 ID（示例中使用时间戳）
	t := time.Now()
	ctt, ss := logger.CreateSpan(ctx, info.FullMethod)
	//c.Request = c.Request.WithContext(ctt)
	defer ss.End()
	// 请求前
	log1 := log.Ctx(ctt)
	log1.Info().Msgf("请求接口 %s", info.FullMethod)

	// 调用后续处理器
	resp, err := handler(ctx, req)

	//// 请求后
	latency := time.Since(t)

	log1.Info().Bool("status", err == nil).
		Dur("latency", latency).
		Msgf("请求接口 %s 结果 %d 耗时 %s", info.FullMethod, err == nil, latency.String())

	return resp, err
}
