package errors

import "github.com/pkg/errors"

var (
	ErrUserExist  = errors.New("user exist")
	ErrUserNotExist = errors.New("User does not exist")
)
