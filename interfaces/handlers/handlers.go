package handlers

import (
	"github.com/gofiber/fiber/v2"
	"massivleads/domain/services"
	smodels "massivleads/prototypes/models"
)

// Handlers contain all handlers for app as method
type Handlers struct {
	services services.Services
}

// SendResponse is a generic method to send http response using current http framework
func (h *Handlers) SendResponse(ctx smodels.RouterCtx, res *smodels.Result) error {
	if res.Exception != nil {
		return ctx.Status(int(res.Exception.StatusCode)).JSON(res.Exception)
	}

	return ctx.Status(fiber.StatusOK).JSON(res.Result)
}

// NewHandlers create a new Handlers instance
func NewHandlers() *Handlers {
	handlers := new(Handlers)
	handlers.services = services.GetServices()

	return handlers
}
