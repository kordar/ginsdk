package validations

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type IValidation interface {
	Tag() string
	ValidationFunc(fl validator.FieldLevel) bool
}

type ValidateContainer struct {
	container map[string]IValidation
}

func NewValidateContainer() *ValidateContainer {
	return &ValidateContainer{
		container: make(map[string]IValidation),
	}
}

func (v *ValidateContainer) AddValidation(validator IValidation) {
	v.container[validator.Tag()] = validator
}

func (v *ValidateContainer) RegisterValidation() {
	if valid, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for tag, validation := range v.container {
			_ = valid.RegisterValidation(tag, validation.ValidationFunc)
		}
	}

}

// InitValidation 初始化自定义验证器
/*
	// ValidateContainer.AddValidation(valid.PhoneValidator{}) // validate phone number
	// ValidateContainer.AddValidation(valid.PasswordValidator{})  // validate password number
*/
func (v *ValidateContainer) InitValidation(validator ...IValidation) *ValidateContainer {
	for _, validation := range validator {
		v.AddValidation(validation)
	}
	return v
}
