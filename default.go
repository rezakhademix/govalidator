package govalidator

// DefaultInt sets a default value for any pointer to an int that is passed
// if value does not exist, will set the default specified as the new value.
//
// Example:
//
//	v := validator.New()
//	var zeroKelvin int64
//	v.DefaultInt(&zeroKelvin, -273)
func (v *Validator) DefaultInt(i *int, val int) *Validator {
	if i == nil || *i == 0 {
		*i = val
	}

	return v
}

// DefaultFloat sets a default value for any pointer to a float that is passed
// if value does not exist, will set the default specified as the new value.
//
// Example:
//
//	v := validator.New()
//	var f float64
//	v.DefaultFloat(&f, 3.14)
func (v *Validator) DefaultFloat(f *float64, val float64) *Validator {
	if f == nil || *f == 0 {
		*f = val
	}

	return v
}

// DefaultString sets a default value for any pointer to a string that is passed.
// if value does not exist, will set the default specified as the new value.
//
// Example:
//
//	v := validator.New()
//	var lang string
//	v.DefaultString(&lang, "persian")
func (v *Validator) DefaultString(s *string, val string) *Validator {
	if s == nil || *s == "" {
		*s = val
	}

	return v
}
