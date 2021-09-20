package middlewares

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"massivleads/domain/entity"
	"massivleads/domain/repository"
	"massivleads/exceptions"
	smodels "massivleads/prototypes/models"
	"massivleads/utils"
)

var userRepo = repository.GetRepositories().UserRepository

type authUser struct{}

// SessionAuth handle auth for session user
func (a *authUser) SessionAuth(token string) (*entity.User, error) {
	at := utils.ValidateJwt(token)
	if at == nil {
		return nil, errors.New("invalid jwt")
	}

	user, err := userRepo.GetUserByUsername(at.Username)

	return user, err
}

//func (a *AuthUser) ApiAuth(token string) (*entity.User, error) {
//
//}

// NewAuthMiddleware create auth handler for routes that required authorization
func NewAuthMiddleware() fiber.Handler {
	auth := new(authUser)

	return func(ctx smodels.RouterCtx) error {
		//token := ctx.Get("access-key")
		//if token != "" && ctx.Locals(AllowedApiKey).(bool) {
		//	// Get api user and return
		//	//ctx.Locals("user", user)
		//
		//	return ctx.Next()
		//}

		token := ctx.Get("session-key")
		if token == "" {
			exp := exceptions.NewBaseException(http.StatusUnauthorized, "Please login to perform this action.")
			return ctx.Status(int(exp.StatusCode)).JSON(exp)
		}

		user, err := auth.SessionAuth(token)

		if err != nil {
			return err
		}

		ctx.Locals("user", user)

		return ctx.Next()
	}
}
