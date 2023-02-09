package validations

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type PasswordValidator struct {
}

func (p PasswordValidator) Tag() string {
	return "password"
}

func (p PasswordValidator) ValidationFunc(fl validator.FieldLevel) bool {
	mobileNum := fl.Field().String()
	fl.StructFieldName()
	regular := "^[a-zA-Z0-9_-]{8,16}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
