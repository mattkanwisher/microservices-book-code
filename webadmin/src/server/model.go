package main

type LoginData struct {
	Validate map[string]string
	Error    string
	Username string
	Password string
}

type RegisterData struct {
	Validate        map[string]string
	Error           string
	Username        string
	Password        string
	ConfirmPassword string
	Email           string
	Name            string
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type UpdateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
}
