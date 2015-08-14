// model.go
package user

import (
	"time"
)

type User struct {
	Id        string    `json:"id" bson:"_id"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"-" bson:"password"`
	Email     string    `json:"email" bson:"email"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type CreateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
