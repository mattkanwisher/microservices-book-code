package main

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrUsernameOrPasswordInvalid = errors.New("Username or password invalid")
	ErrUsernameRequired          = errors.New("Username is required")
	ErrPasswordRequired          = errors.New("Password is required")
	ErrEmailRequired             = errors.New("Email is required")
	ErrNameRequired              = errors.New("Name is required")
	ErrConfirmPasswordRequired   = errors.New("Confirm Password is required")
	ErrEmailInvalid              = errors.New("Email is invalid")
	ErrPasswordMisMatch          = errors.New("Password mismatch")
	ErrOldPasswordRequired       = errors.New("Old Password is required")
)

type errorResponse struct {
	Message interface{} `json:"message"`
}

func ErrorResponse(c *gin.Context, status int, msg interface{}) {
	if msg != nil {
		if err, ok := msg.(error); ok {
			c.JSON(status, &errorResponse{err.Error()})
		} else {
			c.JSON(status, &errorResponse{msg})
		}
	}
}
