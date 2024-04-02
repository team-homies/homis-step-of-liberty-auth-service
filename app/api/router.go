package api

import (
	"main/app/api/patient"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	ptientGroup := app.Group("/api")
	patient.SetApis(ptientGroup)
}
