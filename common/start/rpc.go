package start

import (
	"context"
	"github.com/agamotto-cloud/go-common/common/logger"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"net"
	"time"
)

// RpcServer HttpServer 启动http服务
func RpcServer(srvReg func(srv *grpc.Server)) {
	//设置监听9000端口,添加统一前置处理器
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			otelgrpc.UnaryServerInterceptor(),
			unaryInterceptor,
		),
	}

	srv := grpc.NewServer(opts...)
	srvReg(srv)
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Info().Msg("server start on 9000")
	er := srv.Serve(lis)
	if er != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	t := time.Now()
	ctx, span := logger.CreateSpan(ctx, info.FullMethod)
	defer span.End()

	// 请求前
	logg := log.Ctx(ctx)
	logg.Info().Msgf("请求接口 %s", info.FullMethod)

	// 调用后续处理器
	resp, err := handler(ctx, req)

	//// 请求后
	latency := time.Since(t)

	logg.Info().Bool("status", err == nil).
		Dur("latency", latency).
		Msgf("接口返回 %s 结果 %t 耗时 %s", info.FullMethod, err == nil, latency.String())

	return resp, err
}
