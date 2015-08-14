// service.go
package user

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/satori/go.uuid"
	"time"
)

type Service interface {
	Create(input *CreateInput) (*User, error)
	Get(id string) (*User, error)
	Delete(id string) error
	All() ([]User, error)
	Update(id, name, email string) (*User, error)
	ChangePassword(id, password, oldPassword string) error
	GetByUsernamePassword(username, password string) (*User, error)
}

type UserService struct {
	store Store
}

func New(store Store) Service {
	return &UserService{store}
}

func (s *UserService) Create(input *CreateInput) (*User, error) {
	if v := validateCreate(input); v.HasError() {
		return nil, v.GetError()
	}

	user := &User{
		Id:        uuid.NewV1().String(),
		Username:  input.Username,
		Password:  hashPassword(input.Password),
		Email:     input.Email,
		Name:      input.Name,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.store.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Get(id string) (*User, error) {
	return s.store.GetUser(id)
}

func (s *UserService) Delete(id string) error {
	return s.store.DeleteUser(id)
}

func (s *UserService) All() ([]User, error) {
	return s.store.All()
}

func (s *UserService) Update(id string, name, email string) (*User, error) {
	if v := validateUpdate(id, name, email); v.HasError() {
		return nil, v.GetError()
	}

	return s.store.UpdateUser(id, email, name)
}

func (s *UserService) ChangePassword(id string, password, oldPassword string) error {
	if v := validateChangePassword(id, password, oldPassword); v.HasError() {
		return v.GetError()
	}

	user, err := s.store.GetUser(id)
	if err != nil {
		return err
	}

	hashOldPwd := hashPassword(oldPassword)

	if user.Password != hashOldPwd {
		return ErrOldPasswordInvalid
	}

	return s.store.ChangePassword(id, hashPassword(password))
}

func (s *UserService) GetByUsernamePassword(username, password string) (*User, error) {
	u, err := s.store.GetUserByUsername(username)
	if err != nil {
		return nil, ErrInvalidUsername
	}

	if u.Password != hashPassword(password) {
		return nil, ErrInvalidPassword
	}

	return u, nil
}

func hashPassword(s string) string {
	hasher := sha512.New()
	hasher.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
