package resource

type GetPatientResponse struct {
	Patient PatientResource `json:"patient"`
}

type GetPatientsResponse struct {
	Patients []PatientResource `json:"patients"`
}
