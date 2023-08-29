package rpc

import (
	"buf.build/gen/go/agamotto/test/grpc/go/proto/user/usergrpc"
	"buf.build/gen/go/agamotto/test/protocolbuffers/go/proto/user"
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"time"
)

func Start() {

	srv := grpc.NewServer()
	usergrpc.RegisterUserServiceServer(srv, UserServiceServer{})

}

// 测试user的prc客户端调用,延迟3s后调用
func CallRpc() {

	//延迟3s
	time.Sleep(3 * time.Second)
	//创建RPC的客户端
	conn, _ := grpc.Dial("localhost:9000")
	conn.Connect()
	client := usergrpc.NewUserServiceClient(conn)
	resp, _ := client.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: "test",
	})

	log.Info().Msgf("调用结果 resp: %v", resp)

}
