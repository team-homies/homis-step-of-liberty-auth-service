package resource

type GetPatientsRequest struct {
	Id         uint64 `json:"id" query:"id"`
	Name       string `json:"name" query:"name"`
	MainDoctor string `json:"main_doctor" query:"main_doctor"`
}

type GetPatientRequest struct {
	Id uint64 `json:"id" query:"id"`
}
