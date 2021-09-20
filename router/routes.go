package router

import (
	"github.com/gofiber/fiber/v2"
	"massivleads/domain/entity"
	"massivleads/router/handlers"
	"massivleads/router/middlewares"
	smodels "massivleads/prototypes/models"
)

// CreateRoutes initialize all routes
func CreateRoutes(app smodels.RouterApp) {
	h := handlers.NewHandlers()
	auth := middlewares.NewAuthMiddleware()

	app.Post("/login", h.LoginUser)
	app.Post("/signup", h.RegisterUser)
	app.Get("/refresh", h.RefreshToken)
	app.Get(
		"/test", auth, func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*entity.User)
			return ctx.SendString(user.Username)
		},
	)
}
