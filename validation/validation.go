package validation

import (
	"strings"

	validator "gopkg.in/go-playground/validator.v8"
)

var (
	field         = "{field}"
	param         = "{param}"
	rulesMessages = map[string]map[string]string{
		"string": {
			"min":      "{field} field must be {param} characters at least",
			"max":      "{field} field can be {param} characters at most",
			"required": "{field} field is required!",
			"eqfield":  "{field} field is not equal to {param}",
		},
	}
)

// Messages returns validation errors.
func Messages(errors interface{}) map[string]string {
	messages := make(map[string]string)
	if errs, ok := errors.(validator.ValidationErrors); ok {
		for _, err := range errs {
			messages[strings.ToLower(err.Field)] = addParams(err)
		}
	}

	return messages
}

func addParams(err *validator.FieldError) string {
	str := strings.Fields(rulesMessages[err.Kind.String()][err.Tag])
	for k, v := range str {
		if v == field {
			str[k] = err.Field
		}
		if v == param {
			str[k] = err.Param
		}
	}

	return strings.Join(str, " ")
}
