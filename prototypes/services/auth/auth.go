package auth

import (
	"massivleads/router/dtos"
	"massivleads/prototypes/models"
)

type Service interface {
	RefreshToken(dto dtos.RefreshToken) (*models.Result, error)
}
