package services

import (
	"massivleads/prototypes/models"
	"massivleads/router/dtos"
)

type User interface {
	LoginUser(dto dtos.LoginUser) (*models.Result, error)
	RegisterUser(dto dtos.RegisterUser) (*models.Result, error)
}
