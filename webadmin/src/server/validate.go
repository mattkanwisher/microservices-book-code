package main

import (
	"github.com/plimble/validator"
)

func validateLogin(data *LoginData) validator.Validator {
	v := validator.New()

	v.RequiredString(data.Username, ErrUsernameRequired, "username")
	v.RequiredString(data.Password, ErrPasswordRequired, "password")

	return v
}

func ValidateRegister(data *RegisterData) validator.Validator {
	v := validator.New()

	v.RequiredString(data.Username, ErrUsernameRequired, "username")
	v.RequiredString(data.Password, ErrPasswordRequired, "password")
	v.RequiredString(data.ConfirmPassword, ErrConfirmPasswordRequired, "confirm_password")
	v.Confirm(data.Password, data.ConfirmPassword, ErrPasswordMisMatch, "password")
	v.RequiredString(data.Name, ErrNameRequired, "name")
	v.RequiredString(data.Email, ErrEmailRequired, "email")
	v.Email(data.Email, ErrEmailInvalid, "email")

	return v
}

func validateCreateUser(req *CreateUserRequest) validator.Validator {
	v := validator.New()

	v.RequiredString(req.Username, ErrUsernameRequired, "username")
	v.RequiredString(req.Password, ErrPasswordRequired, "password")
	v.RequiredString(req.Name, ErrNameRequired, "name")
	v.RequiredString(req.Email, ErrEmailRequired, "email")
	v.Email(req.Email, ErrEmailInvalid, "email")

	return v
}

func validateUpdateUser(req *UpdateUserRequest) validator.Validator {
	v := validator.New()

	v.RequiredString(req.Name, ErrNameRequired, "name")
	v.RequiredString(req.Email, ErrEmailRequired, "email")

	return v
}

func validateChangePassword(req *ChangePasswordRequest) validator.Validator {
	v := validator.New()

	v.RequiredString(req.Password, ErrPasswordRequired, "password")
	v.RequiredString(req.OldPassword, ErrOldPasswordRequired, "old_password")

	return v
}
