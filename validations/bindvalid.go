package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// DefaultValidParams 验证结构体，若Valid存在，并验证
func DefaultValidParams(c *gin.Context, stru interface{}) error {
	if err := c.ShouldBind(stru); err != nil {
		return err
	}
	// 获取验证器
	valid, err := GetValidator(c)
	if err != nil {
		return err
	}
	// 获取翻译器
	err = valid.Struct(stru)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return e
		}
	}
	refParams := reflect.ValueOf(stru) // 需要传入指针，后面再解析
	validMethod := refParams.MethodByName("Valid")
	if validMethod.IsValid() {
		v := validMethod.Call(make([]reflect.Value, 0))
		if e := v[0].Interface(); e != nil {
			return e.(error)
		}
	}
	return nil
}

// ValidParamsAndServiceMethod 验证结构体，并验证关联对象的方法
func ValidParamsAndServiceMethod(c *gin.Context, params interface{}, targetService interface{}, methods ...string) error {
	// 获取验证器
	valid, err := GetValidator(c)
	if err != nil {
		return err
	}
	// 获取翻译器
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return e
		}
	}
	if targetService == nil {
		return nil
	}

	refParams := reflect.ValueOf(targetService) // 需要传入指针，后面再解析
	for _, method := range methods {
		validMethod := refParams.MethodByName(method)
		if validMethod.IsValid() {
			P := make([]reflect.Value, 2)
			P[0] = reflect.ValueOf(c)
			P[1] = reflect.ValueOf(params)
			v := validMethod.Call(P)
			if e := v[0].Interface(); e != nil {
				return e.(error)
			}
		}
	}
	return nil
}

func GetValidator(c *gin.Context) (*validator.Validate, error) {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil, errors.New("no validations set")
	}
	return validate, nil
}
