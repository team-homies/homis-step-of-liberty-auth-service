package dosage

import (
	"context"
	"main/app/grpc/proto/dosage"

	"google.golang.org/grpc"
)

type server struct {
	dosage.DosageServiceServer
}

// proto에서 정의한 함수명과 동일하게 할 것
func (s *server) GetDosageById(ctx context.Context, in *dosage.DosageRequest) (*dosage.DosageResponse, error) {
	// 서비스 함수 실행 or 로직 구현
	return &dosage.DosageResponse{
		DosageNo: in.GetDosageNo(),
		DrugInfo: &dosage.Drug{
			Name:        "SA225P2",
			Description: "만병통치약",
			Usage:       "1일 1024회 복용",
			SideEffect:  "72시간 뒤 사망 할 수 있음",
		},
	}, nil
}

func RegisterDosageService(grpcServer *grpc.Server) {
	dosage.RegisterDosageServiceServer(grpcServer, &server{})
}
