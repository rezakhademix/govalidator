package govalidator

// DefaultInt sets a default value for any pointer to an int thats is passed
// if value does not exists, then set the default specified as the new value.
func (v *Validator) DefaultInt(i *int, val int) *Validator {
	if i == nil || *i == 0 {
		*i = val
	}

	return v
}

// DefaultFloat sets a default value for any pointer to an float thats is passed
// if value does not exists, then set the default specified as the new value.
func (v *Validator) DefaultFloat(f *float64, val float64) *Validator {
	if f == nil || *f == 0 {
		*f = val
	}

	return v
}

// DefaultString sets a default value for a pointer to a string.
// if value does not exists, then set the default specified as the new value.
func (v *Validator) DefaultString(s *string, val string) *Validator {
	if s == nil || *s == "" {
		*s = val
	}

	return v
}
