package validator

import (
	"fmt"
	"strings"
)

const (
	// RequiredMethod represents the method name used for finding the default error message.
	RequiredMethod = "required"
	// RequiredErrorMessage is the default error message format for required fields.
	RequiredErrorMessage = "%s is required"
)

// RequiredString checks if a string value is empty or not.
func (v *Validator) RequiredString(value, field string, msg ...string) *Validator {
	if strings.TrimSpace(value) == "" {
		if msg[0] == "" {
			msg[0] = fmt.Sprintf(RequiredErrorMessage, field)
		}

		v.addError(field, msg[0])
	}

	return v
}

// RequiredInt checks if a integer value is provided or not.
func (v *Validator) RequiredInt(value int, field string, msg ...any) *Validator {
	if value == 0 {
		msg := v.errMsg(RequiredMethod, field, msg)

		v.addError(field, msg)
	}

	return v
}
