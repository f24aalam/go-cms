package main

import (
	"go-cms/app/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	v := validator.New()

	app := iris.New()

	app.Validator = v

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
