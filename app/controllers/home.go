package controllers

import "github.com/kataras/iris/v12"

type HomeController struct {
}

func (c *HomeController) Get(ctx iris.Context) {
	data := iris.Map{
		"Title": "Index",
	}

	ctx.ViewLayout("main")
	ctx.View("index", data)
}
