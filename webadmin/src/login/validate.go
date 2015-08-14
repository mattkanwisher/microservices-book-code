package login

import (
	"github.com/plimble/validator"
)

func validateLogin(username, password string) validator.Validator {
	v := validator.New()

	v.RequiredString(username, ErrUsernameRequired)
	v.RequiredString(password, ErrPasswordRequired)

	return v
}
