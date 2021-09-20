package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"massivleads/exceptions"
	"massivleads/interfaces/dtos"
	"massivleads/logger"
	smodels "massivleads/prototypes/models"
	"massivleads/utils"
)

// RefreshToken refresh session token using refresh token
func (h *Handlers) RefreshToken(ctx smodels.RouterCtx) error {
	authService := h.services.Auth
	dto := new(dtos.RefreshToken)

	if err := ctx.QueryParser(dto); err != nil {
		logger.Error("Failed to parse dto", errors.Wrap(err, "handlers.RefreshToken"))
		return err
	}

	if err := utils.ValidateDto(dto); err != nil {
		res := exceptions.NewBaseException(http.StatusBadRequest, "Invalid session.")
		return ctx.Status(fiber.StatusBadRequest).JSON(res)
	}

	res, err := authService.RefreshToken(*dto)

	if err != nil {
		return err
	}

	return h.SendResponse(ctx, res)
}
