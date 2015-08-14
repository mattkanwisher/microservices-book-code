package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"login"
	"register"
	"user"
	userStore "user/store/mongo"
)

func main() {
	// database
	mgoSession, err := mgo.Dial("128.199.130.44:27017")
	if err != nil {
		panic(err)
	}

	// register services
	userService := user.New(userStore.New(mgoSession, "example"))
	loginService := login.New(userService)
	registerService := register.New(userService)

	// create server
	r := render.New(render.Options{
		Directory:     "server/views",
		Layout:        "layout",
		Extensions:    []string{".html"},
		Delims:        render.Delims{"[[", "]]"},
		IsDevelopment: true,
	})

	g := gin.New()
	g.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Content-Type",
		Credentials:     true,
		ValidateHeaders: false,
	}))
	g.Use(sessions.Sessions("webadmin", sessions.NewCookieStore([]byte("something-very-secret"))))
	g.Static("/assets", "./public")

	// register handlers
	NewUserHandlers(userService, r).RegisterHandlers(g)
	NewFrontendHandlers(loginService, registerService, r).RegisterHandlers(g)

	// run server
	g.Run(":3000")
}
