package login

import (
	"user"
)

type Service interface {
	Login(username, password string) (*LoginInfo, error)
}

type LoginService struct {
	userService user.Service
}

func New(us user.Service) Service {
	return &LoginService{us}
}

func (s *LoginService) Login(username, password string) (*LoginInfo, error) {
	if v := validateLogin(username, password); v.HasError() {
		return nil, v.GetError()
	}

	user, err := s.userService.GetByUsernamePassword(username, password)
	if err != nil {
		return nil, ErrUsernameOrPasswordInvalid
	}

	info := &LoginInfo{
		Id: user.Id,
	}

	return info, nil
}
