package validator

import (
	"time"
)

const (
	// Date represents rule name which will be used to find the default error message.
	Date = "date"
	// DateMsg is the default error message format for fields with Date validation rule.
	DateMsg = "%s has wrong date format"
)

// Date checks the value under validation to be a valid, non-relative date.
//
// Example:
//
//	validator.Date("2024-03-09", "2006-01-02", "birthdate", "birthdate must be a valid date in the format YYYY-MM-DD.")
func (v *Validator) Date(d, layout, field, msg string) *Validator {
	_, err := time.Parse(layout, d)
	if err != nil {
		v.Check(false, field, v.msg(Date, msg, field))
	}

	return v
}
