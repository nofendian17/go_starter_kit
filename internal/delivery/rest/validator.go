package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/model/response"
	"reflect"
	"strings"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	return &CustomValidator{
		validator: v,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var errVal response.ErrorValidation
		for _, e := range err.(validator.ValidationErrors) {
			fe := e.Tag()
			if param := e.Param(); len(param) >= 0 {
				fe = fe + ":" + param
			}
			e := response.FieldError{
				Field: e.Field(),
				Error: fe,
			}
			errVal = append(errVal, e)
		}

		return &errVal
	}

	return nil
}

func (cv *CustomValidator) ValidateStruct(i interface{}) error {
	return cv.validator.Struct(i)
}
