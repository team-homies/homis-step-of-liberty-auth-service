package userlist

import (
	"context"
	"main/app/api/common"
	"main/app/grpc/proto/userlist"
	"main/database/repository"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	userlist.UserListServiceServer
}

func (s *server) GetUserList(ctx context.Context, in *userlist.UserListRequest) (*userlist.UserListResponse, error) {
	userlistRepository := repository.NewRepository()

	// 1. 레포지토리 쿼리에서 유저리스트 함수 호출
	res, err := userlistRepository.FindUserList(uint(in.UserId))
	if err != nil {
		return nil, err
	}

	// 2. 수집률을 담기 위해 수집률grpc호출
	per, err := common.GetRateGrpc(uint(in.UserId))
	if err != nil {
		return nil, err
	}

	// 3. 유저코드를 담기 위해 성취도 단계 함수 호출
	code := common.PercentCal(uint64(per))
	visualInfo, err := userlistRepository.FindVisualCode(code)

	return &userlist.UserListResponse{
		UserId:     uint64(res.ID),
		Email:      res.Email,
		Nickname:   res.Nickname,
		Profile:    res.Profile,
		VisualCode: visualInfo.Code,
		Created:    res.CreatedAt.Format(time.RFC3339),
	}, nil
}

func RegisterUserlistService(grpcServer *grpc.Server) {
	userlist.RegisterUserListServiceServer(grpcServer, &server{})
}
