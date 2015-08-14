package user

import "github.com/stretchr/testify/mock"

type MockStore struct {
	mock.Mock
}

func (m *MockStore) GetUser(id string) (*User, error) {
	ret := m.Called(id)

	var r0 *User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*User)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockStore) CreateUser(u *User) error {
	ret := m.Called(u)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) UpdateUser(id string, email string, name string) (*User, error) {
	ret := m.Called(id, email, name)

	var r0 *User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*User)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *MockStore) ChangePassword(password string) error {
	ret := m.Called(password)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) DeleteUser(id string) error {
	ret := m.Called(id)

	r0 := ret.Error(0)

	return r0
}
func (m *MockStore) All() ([]User, error) {
	ret := m.Called()

	var r0 []User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]User)
	}
	r1 := ret.Error(1)

	return r0, r1
}
