package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"massivleads/exceptions"
	"massivleads/interfaces/dtos"
	"massivleads/logger"
	smodels "massivleads/prototypes/models"
	"massivleads/utils"
)

// RegisterUser register new user
func (h *Handlers) RegisterUser(ctx smodels.RouterCtx) error {
	dto := new(dtos.RegisterUser)
	if err := ctx.BodyParser(dto); err != nil {
		logger.Error("Failed to parse dto", errors.Wrap(err, "handlers.RegisterUser"))
		return err
	}

	if err := utils.ValidateDto(dto); err != nil {
		res := new(exceptions.BaseException)

		res.Error = "Bad Request"
		res.StatusCode = fiber.StatusBadRequest
		res.Message = strings.Join(*err, "\n")

		return ctx.Status(fiber.StatusBadRequest).JSON(res)
	}

	userService := h.services.User
	res, err := userService.RegisterUser(*dto)

	if err != nil {
		return err
	}

	return h.SendResponse(ctx, res)
}

// LoginUser login user
func (h *Handlers) LoginUser(ctx smodels.RouterCtx) error {
	dto := new(dtos.LoginUser)
	if err := ctx.BodyParser(dto); err != nil {
		logger.Error("Failed to parse dto", errors.Wrap(err, "handlers.LoginUser"))
		return err
	}

	if err := utils.ValidateDto(dto); err != nil {
		res := new(exceptions.BaseException)

		res.Error = "Bad Request"
		res.StatusCode = fiber.StatusBadRequest
		res.Message = strings.Join(*err, "\n")

		return ctx.Status(fiber.StatusBadRequest).JSON(res)
	}

	userService := h.services.User
	res, err := userService.LoginUser(*dto)

	if err != nil {
		return err
	}

	return h.SendResponse(ctx, res)
}
