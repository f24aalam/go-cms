package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func Validate(ctx iris.Context, form interface{}) {
	err := ctx.ReadForm(form)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.WriteString(err.Error())
			return
		}

		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = ctx.Tr(err.Tag(), strings.ToLower(err.Field()))
		}

		s := sessions.Get(ctx)
		s.SetFlash("errors", errors)
		ctx.Redirect(ctx.GetCurrentRoute().Path())
	}
}
