package grpc

import (
	"main/app/grpc/service/auth"

	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	auth.RegisterAuthService(grpcServer)
}
