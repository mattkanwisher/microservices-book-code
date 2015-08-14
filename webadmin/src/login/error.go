package login

import (
	"errors"
)

var (
	ErrUsernameOrPasswordInvalid = errors.New("Username or password invalid")
	ErrUsernameRequired          = errors.New("Username is required")
	ErrPasswordRequired          = errors.New("Password is required")
)
