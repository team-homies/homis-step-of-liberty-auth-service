package auth

import (
	"main/app/api/auth/handler"
	"main/constant"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetApis(route fiber.Router) {
	h := handler.NewAuthHandler()
	// 토큰 발급
	route.Post(constant.GetPath().Auth.CreateToken, h.CreateToken)

	// 토큰 재발급
	route.Post(constant.GetPath().Auth.UpdateRefreshToken, h.UpdateRefreshToken)

	// 사용자 정보 조회
	route.Get(constant.GetPath().Auth.GetUserInfo, middleware.JWTMiddleware, h.GetUserInfo)

	// 사용자 본인 정보 수정
	route.Put(constant.GetPath().Auth.UpdateUserInfo, middleware.JWTMiddleware, h.UpdateUserInfo)

}
