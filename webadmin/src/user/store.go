// store.go
package user

type Store interface {
	GetUser(id string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(u *User) error
	UpdateUser(id, email, name string) (*User, error)
	ChangePassword(id, password string) error
	DeleteUser(id string) error
	All() ([]User, error)
}
