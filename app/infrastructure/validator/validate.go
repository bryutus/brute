package validator

import (
	"fmt"

	"github.com/bryutus/brute/app/infrastructure/persistence"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var isIso639Alpha2 validator.Func = func(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return iso639Alpha2[val]
}

var existsLanguageCode validator.Func = func(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	r := persistence.NewAphorismPersistence()

	a, err := r.FindBy(val)
	if err != nil {
		return false
	}

	if a != nil {
		return false
	}

	return true
}

var notExistsLanguageCode validator.Func = func(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	r := persistence.NewAphorismPersistence()

	a, err := r.FindBy(val)
	if err != nil {
		return true
	}

	if a == nil {
		return false
	}

	return true
}

func Register() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		return fmt.Errorf("Failed to register custom validation")
	}

	v.RegisterValidation("iso639_1_alpha2", isIso639Alpha2)
	v.RegisterValidation("exists_language_code", existsLanguageCode)
	v.RegisterValidation("not_exists_language_code", notExistsLanguageCode)

	return nil
}
