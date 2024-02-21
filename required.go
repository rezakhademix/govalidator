package validator

import "fmt"

const (
	RequiredMethod       = "required"
	RequiredErrorMessage = "%s is required"
)

func (v *Validator) RequiredString(value, field string, msg ...string) *Validator {
	if value == "" {
		if msg[0] == "" {
			msg[0] = fmt.Sprintf(RequiredErrorMessage, field)
		}

		v.addErrors(field, msg[0])
	}

	return v
}

func (v *Validator) RequiredInt(value int, field string, msgArgs ...any) *Validator {
	if value == 0 {
		msg := FindErrorMessage(RequiredMethod, field, msgArgs)

		v.addErrors(field, msg)
	}

	return v
}
