package handler

import (
	"main/app/api/auth/resource"
	"main/app/api/auth/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	CreateToken(c *fiber.Ctx) error
}

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler() handler {
	return &authHandler{
		service: service.NewAuthService(),
	}
}

func (h *authHandler) CreateToken(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.CreateTokenRequest)
	err := ctx.C.BodyParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusBadRequest)
	}

	res, err := h.service.CreateToken(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
