package grpc

import (
	"main/app/grpc/service/dosage"
	"main/app/grpc/service/patient"

	"google.golang.org/grpc"
)

func RegisterServices(grpcServer *grpc.Server) {
	dosage.RegisterDosageService(grpcServer)
	patient.RegisterPatientService(grpcServer)
}
