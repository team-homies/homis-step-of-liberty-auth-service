package api

import (
	"main/app/api/auth"
	"main/app/api/patient"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	apiGroup := app.Group("/auth")
	patient.SetApis(apiGroup)
	auth.SetApis(apiGroup)
}
