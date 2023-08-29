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

// HttpServer 启动http服务
func RpcServer(srvReg func(srv *grpc.Server)) {
	//设置监听9000端口,添加统一前置处理器
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unaryInterceptor,
			otelgrpc.UnaryServerInterceptor()),
		//grpc.ChainStreamInterceptor(streamInterceptor, otelgrpc.StreamServerInterceptor()),
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

//func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
//	t := time.Now()
//	ctt, span := logger.CreateSpan(ss.Context(), info.FullMethod)
//	defer span.End()
//	// 请求前
//	log1 := log.Ctx(ctt)
//	log1.Info().Msgf("请求接口 %s", info.FullMethod)
//
//	return err
//
//}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// 生成请求的 ID（示例中使用时间戳）
	//如果ctx中有traceid则使用ctx中的traceid
	//logger.CreateSpan()
	t := time.Now()
	ctt, ss := logger.CreateSpan(ctx, info.FullMethod)

	defer ss.End()
	// 请求前
	log1 := log.Ctx(ctt)
	log1.Info().Msgf("请求接口 %s", info.FullMethod)

	// 调用后续处理器
	resp, err := handler(ctt, req)

	//// 请求后
	latency := time.Since(t)

	log1.Info().Bool("status", err == nil).
		Dur("latency", latency).
		Msgf("请求接口 %s 结果 %t 耗时 %s", info.FullMethod, err == nil, latency.String())

	return resp, err
}
