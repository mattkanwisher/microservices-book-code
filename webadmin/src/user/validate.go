// validate.go
package user

import (
	"github.com/plimble/validator"
)

func validateCreate(input *CreateInput) validator.Validator {
	v := validator.New()

	v.RequiredString(input.Username, ErrUsernameRequired)
	v.RequiredString(input.Password, ErrPasswordRequired)
	v.RequiredString(input.Name, ErrNameRequired)
	v.RequiredString(input.Email, ErrEmailRequired)
	v.Email(input.Email, ErrEmailInvalid)

	return v
}

func validateUpdate(id, name, email string) validator.Validator {
	v := validator.New()

	v.RequiredString(id, ErrNameRequired)
	v.RequiredString(name, ErrNameRequired)
	v.RequiredString(email, ErrEmailRequired)

	return v
}

func validateChangePassword(id, password, oldPassword string) validator.Validator {
	v := validator.New()

	v.RequiredString(id, ErrNameRequired)
	v.RequiredString(password, ErrPasswordRequired)
	v.RequiredString(oldPassword, ErrOldPasswordRequired)

	return v
}
