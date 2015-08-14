package register

import (
	"github.com/plimble/validator"
)

func validateRegister(form *RegisterForm) validator.Validator {
	v := validator.New()

	v.RequiredString(form.Username, ErrUsernameRequired)
	v.RequiredString(form.Password, ErrPasswordRequired)
	v.RequiredString(form.Name, ErrNameRequired)
	v.RequiredString(form.Email, ErrEmailRequired)
	v.Email(form.Email, ErrEmailInvalid)

	return v
}
