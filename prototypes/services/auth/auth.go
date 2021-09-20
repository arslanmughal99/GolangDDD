package auth

import (
	"massivleads/interfaces/dtos"
	"massivleads/prototypes/models"
)

type Service interface {
	RefreshToken(dto dtos.RefreshToken) (*models.Result, error)
}
