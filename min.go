package govalidator

import "strings"

const (
	// Min represents rule name which will be used to find the default error message.
	Min = "min"
	// MinString represents the rule name which will be used to find the default error message.
	MinString = "minString"
	// MinMsg is the default error message format for fields with Min validation rule.
	MinMsg = "%s should be more than %v"
	// MinStringMsg is the default error message format for fields with MinString validation rule.
	MinStringMsg = "%s should has more than %v characters"
)

// MinInt checks if the given integer value is greater than or equal the given min value.
//
// Example:
//
//	govalidator.MinInt(18, 0, "age", "age must be at least 0.")
func (v *Validator) MinInt(i, min int, field, msg string) *Validator {
	v.Check(i >= min, field, v.msg(Min, msg, field, min))

	return v
}

// MinFloat checks if the given float value is greater than or equal the given min value.
//
// Example:
//
//	govalidator.MinFloat(5.0, 0.0, "height", "height must be at least 0.0 meters.")
func (v *Validator) MinFloat(f, min float64, field, msg string) *Validator {
	v.Check(f >= min, field, v.msg(Min, msg, field, min))

	return v
}

// MinString checks if the length of given string is greater than or equal the given min value.
//
// Example:
//
//	govalidator.MinString("rey", 5, "name", "name should has more than 5 characters.")
func (v *Validator) MinString(s string, minLen int, field, msg string) *Validator {
	v.Check(len(strings.TrimSpace(s)) >= minLen, field, v.msg(MinString, msg, field, minLen))

	return v
}
