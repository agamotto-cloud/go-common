package rpc

import (
	"buf.build/gen/go/agamotto/test/grpc/go/proto/user/usergrpc"
	"buf.build/gen/go/agamotto/test/protocolbuffers/go/proto/user"
	"context"
	"github.com/agamotto-cloud/go-common/common/logger"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func Start() {

	srv := grpc.NewServer()
	usergrpc.RegisterUserServiceServer(srv, UserServiceServer{})

}

// CallRpc 测试user的prc客户端调用,延迟3s后调用
func CallRpc() {
	ctt, ss := logger.CreateSpan(context.Background(), "test")
	logg := log.Ctx(ctt)
	defer ss.End()

	//延迟3s
	time.Sleep(3 * time.Second)
	//创建RPC的客户端
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))

	if err != nil {
		logg.Fatal().Err(err).Msg("did not connect")
	}
	conn.Connect()
	logg.Info().Msgf("调用参数 ctt: %v", "test")

	client := usergrpc.NewUserServiceClient(conn)
	resp, _ := client.CreateUser(ctt, &user.CreateUserRequest{
		Username: "test",
	})

	logg.Info().Msgf("调用结果 resp: %v", resp)

}

//func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//	//将ctx里的trace传递过去
//	otelgrpc.UnaryClientInterceptor()()
//
//	return cc.Invoke(ctx, method, req, reply, opts...)
//
//}
