package user

import (
	"massivleads/domain/repository"
	"massivleads/domain/services/shared"
)

type Service struct{}

var (
	sessionService = shared.GetSharedServices().Session
	userRepository = repository.GetRepositories().UserRepository
)
