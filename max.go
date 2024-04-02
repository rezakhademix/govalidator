package govalidator

import "strings"

const (
	// Max represents rule name which will be used to find the default error message.
	Max = "max"
	// MaxString represents rule name which will be used to find the default error message.
	MaxString = "maxString"
	// MaxMsg is the default error message format for fields with Max validation rule.
	MaxMsg = "%s should be less than %v"
	// MaxStringMsg is the default error message format for fields with MaxString validation rule.
	MaxStringMsg = "%s should has less than %v characters"
)

// MaxInt checks if the integer value is less than or equal the given max value.
//
// Example:
//
//	validator.MaxInt(10, 100, "age", "age must be less than 100.")
func (v *Validator) MaxInt(i, max int, field, msg string) *Validator {
	v.Check(i <= max, field, v.msg(Max, msg, field, max))

	return v
}

// MaxFloat checks if the given float value is less than or equal the given max value.
//
// Example:
//
//	validator.MaxFloat(3.5, 5.0, "height", "height must be less than 5.0 meters.")
func (v *Validator) MaxFloat(f, max float64, field, msg string) *Validator {
	v.Check(f <= max, field, v.msg(Max, msg, field, max))

	return v
}

// MaxString checks if the length of given string is less than or equal the given max value.
//
// Example:
//
//	validator.MaxString("rey", 5, "name", "name should has less than 5 characters.")
func (v *Validator) MaxString(s string, maxLen int, field, msg string) *Validator {
	v.Check(len(strings.TrimSpace(s)) <= maxLen, field, v.msg(MaxString, msg, field, maxLen))

	return v
}
