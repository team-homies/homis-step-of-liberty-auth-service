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
			Auth: core.AuthPath{
				CreateToken:        "/login",
				UpdateRefreshToken: "/user",
				GetUserInfo:        "/user",
			},
		}
	})

	return instance
}
