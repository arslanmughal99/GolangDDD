package middlewares

import "massivleads/domain/entity"

type AuthUser interface {
	//ApiAuth(token string) (*entity.User, error)
	SessionAuth(token string) (*entity.User, error)
}
