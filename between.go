package govalidator

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	// Between represents rule name which will be used to find the default error message.
	Between = "between"
	// BetweenMsg is the default error message format for fields with Between rule.
	BetweenMsg = "%s should be greater than or equal %v and less than or equal %v"
)

// BetweenInt checks the value under validation to have an integer value between the given min and max.
//
// Example:
//
//	v := validator.New()
//	v.BetweenInt(21, 1, 10, "age", "age must be between 1 and 10.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) BetweenInt(i, min, max int, field, msg string) Validator {
	if i < min || i > max {
		v.addError(field, v.msg(Between, msg, field, min, max))
	}

	return v
}

// BetweenFloat checks the field under validation to have a float value between the given min and max.
//
// Example:
//
//	v := validator.New()
//	v.BetweenFloat(3.5, 2.0, 5.0, "height", "height must be between 2.0 and 5.0 meters.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) BetweenFloat(f, min, max float64, field, msg string) Validator {
	if f < min || f > max {
		v.addError(field, v.msg(Between, msg, field, min, max))
	}

	return v
}

// BetweenString checks if the length of given string to have an integer value between the given min and max.
//
// Example:
//
//	v := validator.New()
//	v.BetweenString("Obi-one", 3, 10, "name", "name must be between 3 and 10 characters.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) BetweenString(s string, minLen, maxLen int, field, msg string) Validator {
	length := utf8.RuneCountInString(strings.TrimSpace(s))
	if length < minLen || length > maxLen {
		v.addError(field, v.msg(Between, msg, fmt.Sprintf("%s length", field), minLen, maxLen))
	}

	return v
}
