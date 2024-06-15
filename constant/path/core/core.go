// 내부 API 경로 정의 패키지
package core

type InternalApi struct {
	Patient PatientPath
	Doctor  DoctorPath
	Auth    AuthPath
}

type PatientPath struct {
	GetPatient    string
	GetPatients   string
	CreatePatient string
	UpdatePatient string
	DeletePatient string
}

type DoctorPath struct {
	GetDoctor     string
	GetPatients   string
	CreatePatient string
	UpdatePatient string
	DeletePatient string
}

type AuthPath struct {
	CreateToken        string
	UpdateRefreshToken string
	GetUserInfo        string
	UpdateUserInfo     string
	FindVisual         string
	FindVisualCode     string
}
