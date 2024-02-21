package validator

import "fmt"

const (
	// RequiredMethod determine method name for finding default error message
	RequiredMethod = "required"

	// RequiredErrorMessage determine method default error message
	RequiredErrorMessage = "%s is required"
)

// RequiredString check if string value is empty return validation error message
func (v *Validator) RequiredString(value, field string, msg ...string) *Validator {
	if value == "" {
		if msg[0] == "" {
			msg[0] = fmt.Sprintf(RequiredErrorMessage, field)
		}

		v.addErrors(field, msg[0])
	}

	return v
}

// RequiredInt check if integer value is empty return validation error message
func (v *Validator) RequiredInt(value int, field string, msgArgs ...any) *Validator {
	if value == 0 {
		msg := FindErrorMessage(RequiredMethod, field, msgArgs)

		v.addErrors(field, msg)
	}

	return v
}
