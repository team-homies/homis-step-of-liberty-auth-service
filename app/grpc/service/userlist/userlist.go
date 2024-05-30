package userlist

import (
	"context"
	"main/app/grpc/proto/userlist"
	"main/constant/common"
	"main/database/repository"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	userlist.UserListServiceServer
}

func (s *server) GetUserList(ctx context.Context, in *userlist.UserListRequest) (*userlist.UserListResponse, error) {
	userlistRepository := repository.NewRepository()
	// 1. 유저리스트 레포지토리
	res, err := userlistRepository.FindUserList(uint(in.UserId))
	if err != nil {
		return nil, err
	}

	// 2. 수집률 담기
	var per uint

	// 3. 유저코드
	code := common.PercentCal(per)

	// 4. 리턴
	return &userlist.UserListResponse{
		UserId:     int64(res.ID),
		Email:      res.Email,
		Nickname:   res.Nickname,
		Profile:    res.Profile,
		VisualCode: code,
		Created:    res.CreatedAt.Format(time.RFC3339),
	}, nil
}

func RegisterUserlistService(grpcServer *grpc.Server) {
	userlist.RegisterUserListServiceServer(grpcServer, &server{})
}
