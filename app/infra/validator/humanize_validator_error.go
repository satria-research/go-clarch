package xvalidator

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Lv1Error struct {
	Param   string
	Message string
}

func TranslateError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}

func GenerateHumanizeError(payload any, err error) []Lv1Error {
	out := make([]Lv1Error, 0)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {

		for _, fe := range ve {

			fieldName := fe.Field()
			field, _ := reflect.TypeOf(payload).Elem().FieldByName(fieldName)
			fieldJSONName, _ := field.Tag.Lookup("json")

			out = append(out, Lv1Error{
				Param:   fieldJSONName,
				Message: TranslateError(fe),
			})
		}
	}

	return out
}
