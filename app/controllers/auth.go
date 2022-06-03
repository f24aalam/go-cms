package controllers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
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
	if s != nil {
		if s.HasFlash() {
			var errors validator.ValidationErrors = (s.GetFlash("errors")).(validator.ValidationErrors)
			fmt.Println(errors[0].Tag())
			if len(errors) > 0 {
				data["Errors"] = errors
			}

		}
	}

	ctx.ViewLayout("auth")
	ctx.View("auth/index", data)
}

func (c *AuthController) PostLogin(ctx iris.Context) {
	var form loginForm
	err := ctx.ReadForm(&form)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.WriteString(err.Error())
			return
		}

		s := sessions.Get(ctx)
		s.SetFlash("errors", err.(validator.ValidationErrors))
		ctx.Redirect(ctx.GetCurrentRoute().Path())
	}
}
