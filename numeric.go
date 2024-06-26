package govalidator

import "strconv"

const (
	// NumericString represents rule name which will be used to find the default error message.
	NumericString = "numericString"
	// NumericStringMsg is the default error message format for fields with NumericString validation rule.
	NumericStringMsg = "%v must be a numeric string"

	// bitSize is a number from 0 to 64
	bitSize = 64
	// base is a number from 2 to 36 or 0
	base = 10
)

// NumericString validates that the field under validation is a numeric string.
//
// Example:
//
//	v := validator.New()
//	v.NumericString("123456789", "postal_code", "postal_code must be a numeric string")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) NumericString(s, field, msg string) Validator {
	_, err := strconv.ParseInt(s, base, bitSize)

	v.check(err == nil, field, v.msg(NumericString, msg, field))

	return v
}
