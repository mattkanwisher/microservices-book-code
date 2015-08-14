// errors.go
package user

import (
	"errors"
)

var (
	ErrIdRequired          = errors.New("Id is required")
	ErrUsernameRequired    = errors.New("Username is required")
	ErrPasswordRequired    = errors.New("Password is required")
	ErrEmailRequired       = errors.New("Email is required")
	ErrNameRequired        = errors.New("Name is required")
	ErrOldPasswordRequired = errors.New("Old Password is required")
	ErrEmailInvalid        = errors.New("Email is invalid")
	ErrOldPasswordInvalid  = errors.New("Old Password is invalid")
	ErrInvalidUsername     = errors.New("Invalid username")
	ErrInvalidPassword     = errors.New("Invalid password")
)
