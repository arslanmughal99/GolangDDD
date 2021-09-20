package user

import (
	"errors"
	"net/http"

	"massivleads/domain/entity"
	"massivleads/exceptions"
	"massivleads/interfaces/dtos"
	"massivleads/logger"
	serrors "massivleads/prototypes/errors"
	"massivleads/prototypes/models"
	"massivleads/utils"
)

// RegisterUser register a nwe user
func (s *Service) RegisterUser(dto dtos.RegisterUser) (*models.Result, error) {

	result := new(models.Result)
	newUser := new(entity.User)
	response := new(dtos.RegisterUserResponse)
	hash, err := utils.HashPassword(dto.Password)

	if err != nil {
		logger.Error("Failed to hash password", err)
		return nil, err
	}

	newUser.Hash = *hash
	newUser.Name = dto.Name
	newUser.Email = dto.Email
	newUser.Username = dto.Username

	err = userRepository.CreateUser(*newUser)

	if errors.Is(err, serrors.ErrUserExist) {
		result.Exception = exceptions.NewBaseException(http.StatusConflict, "Username or email already exist.")

		return result, nil
	}

	if err != nil {
		return nil, err
	}

	response.Username = dto.Username
	result.Result = response

	return result, nil
}
