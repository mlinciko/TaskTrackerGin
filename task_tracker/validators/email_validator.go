package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var EmailValidator validator.Func = func(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if ok {
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		return re.MatchString(email)

	}
	return true
}
