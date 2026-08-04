package govalidator

// DefaultInt sets the given default value on the passed int pointer when it
// currently points to a zero value. A nil pointer is a no-op: there is no
// variable to write the default into.
//
// Example:
//
//	v := validator.New()
//	var zeroKelvin int64
//	v.DefaultInt(&zeroKelvin, -273)
func (v Validator) DefaultInt(i *int, val int) Validator {
	if i != nil && *i == 0 {
		*i = val
	}

	return v
}

// DefaultFloat sets the given default value on the passed float pointer when
// it currently points to a zero value. A nil pointer is a no-op: there is no
// variable to write the default into.
//
// Example:
//
//	v := validator.New()
//	var f float64
//	v.DefaultFloat(&f, 3.14)
func (v Validator) DefaultFloat(f *float64, val float64) Validator {
	if f != nil && *f == 0 {
		*f = val
	}

	return v
}

// DefaultString sets the given default value on the passed string pointer when
// it currently points to an empty string. A nil pointer is a no-op: there is
// no variable to write the default into.
//
// Example:
//
//	v := validator.New()
//	var lang string
//	v.DefaultString(&lang, "persian")
func (v Validator) DefaultString(s *string, val string) Validator {
	if s != nil && *s == "" {
		*s = val
	}

	return v
}
