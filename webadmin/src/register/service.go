package register

import (
	"user"
)

type Service interface {
	Register(form *RegisterForm) error
}

type RegisterService struct {
	userService user.Service
}

func New(us user.Service) Service {
	return &RegisterService{us}
}

func (s *RegisterService) Register(form *RegisterForm) error {
	if v := validateRegister(form); v.HasError() {
		return v.GetError()
	}

	_, err := s.userService.Create(&user.CreateInput{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Name:     form.Name,
	})

	if err != nil {
		return err
	}

	return nil
}
