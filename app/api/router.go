package api

import (
	"main/app/api/auth"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	apiGroup := app.Group("/auth")
	auth.SetApis(apiGroup)
}
