package resource

type PatientResource struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	MainDoctor string `json:"main_doctor"`
}
