package main

import (
	"go-cms/app/controllers"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.Blocks("./resources/views", ".html").Reload(true))
	app.HandleDir("/", iris.Dir("./public"))

	setupRoutes(app)

	app.Listen(":8000")
}

func setupRoutes(app *iris.Application) {
	app.Get("/", new(controllers.HomeController).Index)
}
