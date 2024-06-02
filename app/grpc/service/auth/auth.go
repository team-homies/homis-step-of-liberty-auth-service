package auth

import (
	"context"
	"main/app/grpc/proto/auth"
	"main/middleware"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	auth.AuthServiceServer
}

func RegisterAuthService(grpcServer *grpc.Server) {
	auth.RegisterAuthServiceServer(grpcServer, &server{})
}

func (s *server) GetAuthVerification(ctx context.Context, in *auth.AuthRequest) (*auth.AuthResponse, error) {
	// 서비스 함수 실행 or 로직 구현
	checked, userId, err := middleware.JwtChecker(in.Token)
	var message string
	if err != nil {
		message = err.Error()
	}
	userIddduni, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		message = err.Error()
	}
	return &auth.AuthResponse{
		UserId:  userIddduni,
		IsAuth:  checked,
		Message: message,
	}, err
}
