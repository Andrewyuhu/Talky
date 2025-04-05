package validator

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type validator struct {
	InvalidProperties map[string]string
}

func NewValidator() *validator {
	return &validator{
		InvalidProperties: make(map[string]string),
	}
}

func (v *validator) IsValid() bool {
	return len(v.InvalidProperties) == 0
}

func (v *validator) addError(key string, value string) {
	if _, ok := v.InvalidProperties[key]; !ok {
		v.InvalidProperties[key] = value
	}
}

func MinLength(key string, value string, min int, v *validator) {
	if utf8.RuneCountInString(value) < min {
		v.addError(key, fmt.Sprintf("property must be of at least length %d", min))
	}
}

func MaxLength(key string, value string, max int, v *validator) {
	if utf8.RuneCountInString(value) > max {
		v.addError(key, fmt.Sprintf("property must be less than length %d", max))
	}
}

func ValidEmail(key string, value string, v *validator) {
	if !emailRegex.Match([]byte(value)) {
		v.addError(key, "incorrect email format")
	}
}
