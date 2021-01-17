package validator

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var isIso639Alpha2 validator.Func = func(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return iso639Alpha2[val]
}

func Register() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		return fmt.Errorf("Failed to register custom validation")
	}

	v.RegisterValidation("iso639_1_alpha2", isIso639Alpha2)

	return nil
}
