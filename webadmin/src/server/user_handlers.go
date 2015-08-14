package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"user"
)

type UserHandlers struct {
	userService user.Service
	render      *render.Render
}

func NewUserHandlers(userService user.Service, r *render.Render) *UserHandlers {
	return &UserHandlers{userService, r}
}

func (h *UserHandlers) RegisterHandlers(g *gin.Engine) {
	g.GET("/user", h.All)
	g.GET("/user/:id", h.GetUser)
	g.PUT("/user", h.CreateUser)
	g.POST("/user/:id", h.UpdateUser)
	g.DELETE("/user/:id", h.DeleteUser)
	g.POST("/user/:id/password", h.ChangePassword)
}

func (h *UserHandlers) CreateUser(c *gin.Context) {
	req := &CreateUserRequest{}
	c.BindJSON(req)

	if v := validateCreateUser(req); v.HasError() {
		ErrorResponse(c, 400, v.Messages())
		return
	}

	input := &user.CreateInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Name:     req.Name,
	}

	u, err := h.userService.Create(input)
	if err != nil {
		ErrorResponse(c, 500, err)
		return
	}

	c.JSON(200, u)
}

func (h *UserHandlers) GetUser(c *gin.Context) {
	id := c.Param("id")

	u, err := h.userService.Get(id)
	if err != nil {
		ErrorResponse(c, 500, err)
		return
	}

	c.JSON(200, u)
}

func (h *UserHandlers) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.userService.Delete(id); err != nil {
		ErrorResponse(c, 500, err)
		return
	}

	c.JSON(200, nil)
}

func (h *UserHandlers) All(c *gin.Context) {
	users, err := h.userService.All()
	if err != nil {
		ErrorResponse(c, 500, err)
		return
	}

	c.JSON(200, users)
}

func (h *UserHandlers) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	req := &UpdateUserRequest{}
	c.BindJSON(req)

	if v := validateUpdateUser(req); v.HasError() {
		ErrorResponse(c, 400, v.Messages())
		return
	}

	u, err := h.userService.Update(id, req.Name, req.Email)
	if err != nil {
		ErrorResponse(c, 500, err)
		return
	}

	c.JSON(200, u)
}

func (h *UserHandlers) ChangePassword(c *gin.Context) {
	id := c.Param("id")

	req := &ChangePasswordRequest{}
	c.BindJSON(req)

	if v := validateChangePassword(req); v.HasError() {
		ErrorResponse(c, 400, v.Messages())
		return
	}

	if err := h.userService.ChangePassword(id, req.Password, req.OldPassword); err != nil {
		ErrorResponse(c, 500, err)
		return
	}

	c.JSON(200, nil)
}
