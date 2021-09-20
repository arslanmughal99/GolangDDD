package middlewares

import smodels "massivleads/prototypes/models"

const AllowedApiKey = "ALLOWED_API"

// AllowedApi allow a route for API user
func AllowedApi(ctx smodels.RouterCtx) error {
	ctx.Locals(AllowedApiKey, true)
	return ctx.Next()
}
