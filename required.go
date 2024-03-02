package validator

import "strings"

const (
	// Required represents the rule name which will be used to find the default error message.
	Required = "required"
	// RequiredMsg is the default error message format for required fields.
	RequiredMsg = "%s is required"
)

// RequiredString checks if a string value is empty or not.
func (v *Validator) RequiredString(s, field string, msg string) *Validator {
	v.Check(strings.TrimSpace(s) != "", field, v.msg(Required, field, msg))

	return v
}

// RequiredInt checks if an integer value is provided or not.
func (v *Validator) RequiredInt(i int, field string, msg string) *Validator {
	v.Check(i == 0, field, v.msg(Required, field, msg))

	return v
}
