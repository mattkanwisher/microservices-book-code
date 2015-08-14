package register

import (
	"errors"
)

var (
	ErrUsernameRequired = errors.New("Username is required")
	ErrPasswordRequired = errors.New("Password is required")
	ErrEmailRequired    = errors.New("Email is required")
	ErrNameRequired     = errors.New("Name is required")
	ErrEmailInvalid     = errors.New("Email is invalid")
)
