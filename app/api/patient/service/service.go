package service

import "main/app/api/patient/resource"

type PatientService interface {
	GetPatient(req *resource.GetPatientRequest) (res *resource.GetPatientResponse, err error)
	GetPatients() (resource.GetPatientsResponse, error)
	// CreatePatient(resource.GetPatientRequest) (PatientResource, error)
	// UpdatePatient(id int, patient PatientResource) (PatientResource, error)
	DeletePatient(id int) error
}

func NewPatientService() PatientService {
	return &patientService{
		PatientService: &patientService{},
	}
}

type patientService struct {
	PatientService
}

func (s *patientService) GetPatient(req *resource.GetPatientRequest) (res *resource.GetPatientResponse, err error) {
	res = new(resource.GetPatientResponse)
	res = &resource.GetPatientResponse{
		Patient: resource.PatientResource{
			Id:   1,
			Name: "John Doe",
			Age:  25,
		},
	}
	return
}

func (p *patientService) GetPatients() (resource.GetPatientsResponse, error) {
	return resource.GetPatientsResponse{
		Patients: []resource.PatientResource{
			{
				Id:   1,
				Name: "Will Smith",
				Age:  55,
			},
			{
				Id:   2,
				Name: "Jean Valjean",
				Age:  254,
			},
		},
	}, nil
}
