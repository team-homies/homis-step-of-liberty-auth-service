package constant

import (
	"main/constant/path/core"
	"sync"
)

var (
	instance *core.InternalApi
	once     sync.Once
)

func GetPath() *core.InternalApi {
	once.Do(func() {
		instance = &core.InternalApi{
			Patient: core.PatientPath{
				GetPatient:    "/patient",
				GetPatients:   "/patient/list",
				CreatePatient: "/patient",
				UpdatePatient: "/patient",
				DeletePatient: "/patient",
			},
			Auth: core.AuthPath{
				CreateToken:        "/login",
				UpdateRefreshToken: "/user",
			},
		}
	})

	return instance
}
