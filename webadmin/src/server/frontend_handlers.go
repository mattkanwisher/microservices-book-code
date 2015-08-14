package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"login"
	"register"
)

type FrontendHandlers struct {
	loginService    login.Service
	registerService register.Service
	render          *render.Render
}

func NewFrontendHandlers(
	loginService login.Service,
	registerService register.Service,
	r *render.Render,
) *FrontendHandlers {
	return &FrontendHandlers{
		loginService,
		registerService,
		r,
	}
}

func (h *FrontendHandlers) RegisterHandlers(g *gin.Engine) {
	g.GET("/logout", h.Logout)
	g.GET("/login", h.GetLogin)
	g.POST("/login", h.PostLogin)
	g.GET("/register", h.GetRegister)
	g.POST("/register", h.PostRegister)
	g.GET("/", h.Home)
}

func (h *FrontendHandlers) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(302, "/login")
}

func (h *FrontendHandlers) GetLogin(c *gin.Context) {
	h.render.HTML(c.Writer, 200, "login", nil)
}

func (h *FrontendHandlers) PostLogin(c *gin.Context) {
	data := &LoginData{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	if v := validateLogin(data); v.HasError() {
		data.Validate = v.Messages()
		h.render.HTML(c.Writer, 200, "login", data)
		return
	}

	info, err := h.loginService.Login(data.Username, data.Password)
	if err != nil {
		data.Error = err.Error()
		h.render.HTML(c.Writer, 200, "login", data)
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", info.Id)
	session.Save()

	c.Redirect(302, "/")
}

func (h *FrontendHandlers) GetRegister(c *gin.Context) {
	h.render.HTML(c.Writer, 200, "register", nil)
}

func (h *FrontendHandlers) PostRegister(c *gin.Context) {
	data := &RegisterData{
		Username:        c.PostForm("username"),
		Password:        c.PostForm("password"),
		ConfirmPassword: c.PostForm("confirmpassword"),
		Email:           c.PostForm("email"),
		Name:            c.PostForm("name"),
	}

	if v := ValidateRegister(data); v.HasError() {
		data.Validate = v.Messages()
		h.render.HTML(c.Writer, 200, "register", data)
		return
	}

	form := &register.RegisterForm{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
		Name:     data.Name,
	}

	if err := h.registerService.Register(form); err != nil {
		data.Error = err.Error()
		h.render.HTML(c.Writer, 200, "register", data)
		return
	}

	c.Redirect(302, "/login")
}

func (h *FrontendHandlers) Home(c *gin.Context) {
	session := sessions.Default(c)
	userID, ok := session.Get("user_id").(string)
	if !ok || userID == "" {
		c.Redirect(302, "/login")
		return
	}

	h.render.HTML(c.Writer, 200, "home", nil)
}
