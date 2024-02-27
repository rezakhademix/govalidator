package validator

import "strings"

const (
	// Required represents the rule name which will be used to find the default error message.
	Required = "required"
	// RequiredMsg is the default error message format for required fields.
	RequiredMsg = "%s is required"
)

// RequiredString checks if a string value is empty or not.
func (v *Validator) RequiredString(value, field string, msg string) *Validator {
	if strings.TrimSpace(value) == "" {
		v.addError(field, v.msg(Required, field, msg))
	}

	return v
}

// RequiredInt checks if a integer value is provided or not.
func (v *Validator) RequiredInt(value int, field string, msg string) *Validator {
	if value == 0 {
		v.addError(field, v.msg(Required, field, msg))
	}

	return v
}
