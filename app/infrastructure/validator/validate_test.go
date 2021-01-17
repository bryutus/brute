package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestIsIso639Alpha2(t *testing.T) {
	testCases := []struct {
		in   string
		want bool
	}{
		{"is", true},
		{"isl", false},
		{"qq", false},
	}

	validate := validator.New()
	validate.RegisterValidation("iso639_1_alpha2", isIso639Alpha2)

	for i, test := range testCases {
		errs := validate.Var(test.in, "iso639_1_alpha2")
		if test.want {
			if errs != nil {
				t.Fatalf("Index: %d iso639_1_alpha2 failed Error: %s", i, errs)
			}
		} else {
			if errs == nil {
				t.Fatalf("Index: %d iso639_1_alpha2 failed Error: %s", i, errs)
			}
		}
	}
}
