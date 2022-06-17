package main

import (
	"go-cms/app/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	v := validator.New()
	s := sessions.New(sessions.Config{Cookie: "go-cms-sessions-cookie"})

	app := iris.New()

	app.I18n.Load("./resources/locales/*/*", "en")

	app.Validator = v

	app.Use(s.Handler())
	app.RegisterView(iris.Blocks("./resources/views", ".html").Reload(true))
	app.HandleDir("/", iris.Dir("./public"))

	setupRoutes(app)

	app.Listen(":8000")
}

func setupRoutes(app *iris.Application) {
	m := mvc.New(app.Party("/"))
	m.Handle(new(controllers.HomeController))
	m.Handle(new(controllers.AuthController))
}
