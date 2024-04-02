package grpc

import (
	"context"
	dosage "main/app/grpc/proto/dosage"
	patient "main/app/grpc/proto/patient"

	"google.golang.org/grpc"
)

type GrpcService interface {
	GetPatientById(ctx context.Context, in *patient.PatientRequest) (*patient.PatientResponse, error)
	GetDosageById(ctx context.Context, in *dosage.DosageRequest) (*dosage.DosageResponse, error)
}

type grpcServer struct {
	patientService patient.UnimplementedPatientServiceServer
	dosageService  dosage.UnimplementedDosageServiceServer
}

func NewGrpcServer() GrpcService {
	return &grpcServer{
		patient.UnimplementedPatientServiceServer{},
		dosage.UnimplementedDosageServiceServer{},
	}
}

func ListenGrpcServer() *grpc.Server {
	// 서비스 등록
	s := grpc.NewServer()
	patient.RegisterPatientServiceServer(s, patient.UnimplementedPatientServiceServer{})
	dosage.RegisterDosageServiceServer(s, dosage.UnimplementedDosageServiceServer{})
	return s
}

func (s *grpcServer) GetPatientById(ctx context.Context, in *patient.PatientRequest) (*patient.PatientResponse, error) {
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

func (s *grpcServer) GetDosageById(ctx context.Context, in *dosage.DosageRequest) (*dosage.DosageResponse, error) {
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
