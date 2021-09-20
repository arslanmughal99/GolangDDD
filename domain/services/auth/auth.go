package auth

import (
	"net/http"

	"massivleads/domain/repository"
	"massivleads/domain/services/shared"
	"massivleads/exceptions"
	"massivleads/router/dtos"
	"massivleads/prototypes/models"
	"massivleads/utils"
)

type Service struct{}

func (s *Service) RefreshToken(dto dtos.RefreshToken) (*models.Result, error) {
	sessionService := shared.GetSharedServices().Session
	userRepository := repository.GetRepositories().UserRepository

	result := new(models.Result)
	response := new(dtos.RefreshTokenResponse)
	username := sessionService.GetSession(dto.Refresh)

	if username == nil {
		result.Exception = exceptions.NewBaseException(http.StatusNotFound, "Invalid session.")
		return result, nil
	}

	user, err := userRepository.GetUserByUsername(*username)

	if err != nil {
		return nil, err
	}

	t, err := utils.CreateJwt(user.Username, user.Verified)

	if err != nil {
		return nil, err
	}

	response.Token = *t
	result.Result = response

	return result, nil
}
