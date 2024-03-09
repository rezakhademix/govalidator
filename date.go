package validator

import (
	"time"
)

const (
	// Date represents the rule name which will be used to find the default error message.
	Date = "date"
	// DateMsg is the default error message format for fields with the Date validation rule.
	DateMsg = "%s has wrong date format"
)

// Date checks the field under validation to be a valid, non-relative date.
func (v *Validator) Date(d, layout, field, msg string) *Validator {
	_, err := time.Parse(layout, d)
	if err != nil {
		v.Check(false, field, v.msg(Date, msg, field))
	}

	return v
}
