package auth

import (
	"main/app/api/auth/handler"
	"main/constant"

	"github.com/gofiber/fiber/v2"
)

func SetApis(route fiber.Router) {
	h := handler.NewAuthHandler()
	// 토큰 발급
	route.Post(constant.GetPath().Auth.CreateToken, h.CreateToken)

	// 시각적 성취도 조회
	// route.Get(constant.GetPath().Auth.GetVisualAchievement, middleware.JWTMiddleware, h.GetVisualAchievement)

}
