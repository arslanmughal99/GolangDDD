package middlewares

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	smodels "massivleads/prototypes/models"
)

// SetMiddlewares set middlewares on http app instance
func SetMiddlewares(app smodels.RouterApp) {
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(csrf.New())
}
