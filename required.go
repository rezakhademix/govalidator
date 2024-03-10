package govalidator

import "strings"

const (
	// Required represents rule name which will be used to find the default error message.
	Required = "required"
	// RequiredMsg is the default error message format for fields with required validation rule.
	RequiredMsg = "%s is required"
)

// RequiredString checks if a string value is empty or not.
//
// Example:
//
//	govalidator.RequiredString("hello", "username", "username is required.")
func (v *Validator) RequiredString(s, field string, msg string) *Validator {
	v.Check(strings.TrimSpace(s) != "", field, v.msg(Required, msg, field))

	return v
}

// RequiredInt checks if an integer value is provided or not.
//
// Example:
//
//	govalidator.RequiredInt(42, "age", "age is required.")
func (v *Validator) RequiredInt(i int, field string, msg string) *Validator {
	v.Check(i != 0, field, v.msg(Required, msg, field))

	return v
}

// RequiredSlice checks if a slice has any value or not.
//
// Example:
//
//	govalidator.RequiredSlice([]string{"apple", "banana", "orange"}, "fruits", "at least one fruit must be provided.")
func (v *Validator) RequiredSlice(s []any, field string, msg string) *Validator {
	v.Check(len(s) > 0, field, v.msg(Required, msg, field))

	return v
}

// RequiredFloat checks if a float value is provided or not.
//
// Example:
//
//	govalidator.RequiredFloat(3.5, "weight", "weight is required.")
func (v *Validator) RequiredFloat(f float64, field string, msg string) *Validator {
	v.Check(f != 0.0, field, v.msg(Required, msg, field))

	return v
}
