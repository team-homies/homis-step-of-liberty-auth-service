package handler

import (
	"main/app/api/patient/resource"
	"main/app/api/patient/service"
	"main/common/fiberkit"

	"github.com/gofiber/fiber/v2"
)

type handler interface {
	GetPatient(c *fiber.Ctx) error
	GetPatients(c *fiber.Ctx) error
	CreatePatient(c *fiber.Ctx) error
	UpdatePatient(c *fiber.Ctx) error
	DeletePatient(c *fiber.Ctx) error
}

type patientHandler struct {
	service service.PatientService
}

func NewPatientHandler() handler {
	return &patientHandler{
		service: service.NewPatientService(),
	}
}

func (h *patientHandler) GetPatient(c *fiber.Ctx) error {
	ctx := fiberkit.FiberKit{C: c}
	req := new(resource.GetPatientRequest)
	ctx.C.QueryParser(req)

	res, err := h.service.GetPatient(req)
	if err != nil {
		return ctx.HttpFail(err.Error(), fiber.StatusNotFound)
	}
	return ctx.HttpOK(res)
}

func (h *patientHandler) GetPatients(c *fiber.Ctx) error {
	return nil
}

func (h *patientHandler) CreatePatient(c *fiber.Ctx) error {
	return nil
}

func (h *patientHandler) UpdatePatient(c *fiber.Ctx) error {
	return nil
}

func (h *patientHandler) DeletePatient(c *fiber.Ctx) error {
	return nil
}
