package user

import (
	"errors"
	"net/http"

	"massivleads/exceptions"
	"massivleads/router/dtos"
	"massivleads/logger"
	serrors "massivleads/prototypes/errors"
	"massivleads/prototypes/models"
	"massivleads/utils"
)

// LoginUser login user and return jwt-token
func (s *Service) LoginUser(dto dtos.LoginUser) (*models.Result, error) {
	result := new(models.Result)
	response := new(dtos.LoginUserResponse)

	user, err := userRepository.GetUserByUsernameOrEmail(dto.Username)

	if errors.Is(err, serrors.ErrUserNotExist) {
		result.Exception = exceptions.NewBaseException(http.StatusNotFound, "User not found.")
		return result, nil
	}

	if err != nil {
		return nil, err
	}

	if !utils.ComparePassword(user.Hash, dto.Password) {
		result.Exception = exceptions.NewBaseException(http.StatusForbidden, "Incorrect password.")
		return result, nil
	}

	accessToken, err := utils.CreateJwt(user.Username, user.Verified)

	if err != nil {
		logger.Error("Failed to create access token", err)
		return nil, err
	}

	if dto.Remember {
		rToken, err := sessionService.SetSession(dto.Username)

		if err != nil {
			return nil, err
		}
		response.Refresh = rToken
	}

	response.Token = *accessToken
	result.Result = response

	return result, nil
}
