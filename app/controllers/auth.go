package controllers

import (
	"go-cms/app/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type loginForm struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type AuthController struct{}

func (c *AuthController) GetLogin(ctx iris.Context) {
	data := iris.Map{
		"Title": "Login",
	}

	s := sessions.Get(ctx)
	if s != nil && s.HasFlash() {
		errors := (s.GetFlash("errors")).(map[string]string)
		if len(errors) > 0 {
			data["Errors"] = errors
		}
	}

	ctx.ViewLayout("auth")
	ctx.View("auth/index", data)
}

func (c *AuthController) PostLogin(ctx iris.Context) {
	var form loginForm

	utils.Validate(ctx, &form)
}
