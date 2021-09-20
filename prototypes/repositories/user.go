package repositories

import "massivleads/domain/entity"

type UserRepository interface {
	// GetUserById Get single user from Db by id
	GetUserById(id string) (*entity.User, error)
	// GetUserByUsernameOrEmail Get single user from Db by username or email
	GetUserByUsernameOrEmail(username string) (*entity.User, error)
	// GetUserByUsername get single user by username from Db
	GetUserByUsername(username string) (*entity.User, error)
	// UpdateUser Update single user in Db
	UpdateUser(user entity.User) error
	// CreateUser Create user in Db
	CreateUser(user entity.User) error
}
