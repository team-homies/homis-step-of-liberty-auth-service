package patient

import (
	"main/app/api/patient/handler"
	"main/constant"

	"github.com/gofiber/fiber/v2"
)

func SetApis(route fiber.Router) {
	h := handler.NewPatientHandler()
	// 환자 리스트 조회
	route.Get(constant.GetPath().Patient.GetPatient, h.GetPatient)
	// // 환자 상세 조회
	// route.Get(constant.GetPath().Patient.GetPatients)
	// // 환자 정보 생성
	// route.Post(constant.GetPath().Patient.CreatePatient)
	// // 환자 정보 수정
	// route.Put(constant.GetPath().Patient.UpdatePatient)
	// // 환자 제거
	// route.Delete(constant.GetPath().Patient.DeletePatient)
}
