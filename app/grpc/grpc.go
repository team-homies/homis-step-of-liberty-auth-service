package grpc

import (
	"main/app/grpc/service/userlist"

	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	userlist.RegisterUserlistService(grpcServer)
}
