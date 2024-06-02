package handler

import (
	"main/app/api/auth/resource"
	"main/app/api/auth/service"
	"main/common/fiberkit"
	"main/constant/common"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	CreateToken(c *fiber.Ctx) error
	UpdateRefreshToken(c *fiber.Ctx) error
	GetUserInfo(c *fiber.Ctx) error
	UpdateUserInfo(c *fiber.Ctx) (err error)
	FindVisual(c *fiber.Ctx) error
	FindVisualCode(c *fiber.Ctx) error
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

func (h *authHandler) UpdateRefreshToken(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.UpdateTokenRequest)
	err := ctx.C.BodyParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusBadRequest)
	}

	res, err := h.service.UpdateRefreshToken(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}

func (h *authHandler) GetUserInfo(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	userId := ctx.GetLocalsInt(common.LOCALS_USER_ID)
	res, err := h.service.UserInfo(uint(userId))
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}

// 사용자 본인 정보 수정 body : Nickname, Propile
func (h *authHandler) UpdateUserInfo(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}

	req := new(resource.UpdateUserInfoRequest)
	err := ctx.C.BodyParser(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}

	// 1. userId값 받아오기
	req.Id = uint(ctx.GetLocalsInt(common.LOCALS_USER_ID))

	// 2. 서비스 함수 실행
	err = h.service.UpdateUserInfo(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(err)
}

// 시각적 성취도 조회
func (h *authHandler) FindVisual(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}

	// 1. Id값 받아오기
	userId := ctx.GetLocalsInt(common.LOCALS_USER_ID)

	// 2. 서비스함수 실행
	res, err := h.service.FindVisual(uint(userId))
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}

// 시각적 성취도 코드 조회
func (h *authHandler) FindVisualCode(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}

	// 1. Id값 받아오기
	userId := ctx.GetLocalsInt(common.LOCALS_USER_ID)

	// 2. 서비스함수 실행
	res, err := h.service.FindVisualCode(uint(userId))
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}
