package patient

import (
	"context"
	"main/app/grpc/proto/patient"

	"google.golang.org/grpc"
)

type server struct {
	patient.PatientServiceServer
}

func (s *server) GetPatientById(ctx context.Context, in *patient.PatientRequest) (*patient.PatientResponse, error) {
	// 서비스 함수 실행 or 로직 구현
	return &patient.PatientResponse{
		PatientNo: in.GetPatientNo(),
		Name:      "안녕디지몬",
		Age:       24,
		BirthDate: "2000-09-07",
		Gender:    "unknown",
		Alived:    false,
	}, nil
}

func RegisterPatientService(grpcServer *grpc.Server) {
	patient.RegisterPatientServiceServer(grpcServer, &server{})
}
