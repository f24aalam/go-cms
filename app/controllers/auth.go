package controllers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
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

	ctx.ViewLayout("auth")
	ctx.View("auth/index", data)
}

func (c *AuthController) PostLogin(ctx iris.Context) {
	var form loginForm
	err := ctx.ReadForm(&form)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}

		ctx.StatusCode(iris.StatusBadRequest)
		fmt.Println(err.(validator.ValidationErrors))
	}

	fmt.Println(form)
}
