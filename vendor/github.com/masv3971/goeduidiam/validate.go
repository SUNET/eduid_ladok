package goeduidiam

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func validate(s interface{}) error {
	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("ERR: Field %q of type %q violates rule: %q\n", err.Namespace(), err.Kind(), err.Tag())
		}
		return errors.New("Validation error")
	}
	return nil
}
