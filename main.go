package main

import (
	"ginshop/models"
	"ginshop/routers"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Str2Html":   models.Str2Html,
	})

	r.LoadHTMLGlob("templates/**/*")

	r.Static("/static", "./static")

	store := cookie.NewStore([]byte("secret123456"))

	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRouters(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
