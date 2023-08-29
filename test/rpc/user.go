package rpc

import (
	"buf.build/gen/go/agamotto/test/protocolbuffers/go/proto/user"
	"context"
	"github.com/rs/zerolog/log"
)

type UserServiceServer struct {
}

func (u UserServiceServer) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	//TODO implement me
	logger := log.Ctx(ctx)
	logger.Info().Msgf("CreateUser request: %v", request)
	return &user.CreateUserResponse{
		UserId: 1,
	}, nil

}

func (u UserServiceServer) GetUser(ctx context.Context, request *user.GetUserRequest) (*user.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceServer) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
