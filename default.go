package govalidator

// Sets a default value for any pointer to an int thats is passed
// if value does not exists, then set the default specified as the new value.
func (v Validator) DefaultInt(i *int, val int) {
	if i == nil || *i == 0 {
		*i = val
	}
}

// Sets a default value for any pointer to an float thats is passed
// if value does not exists, then set the default specified as the new value.
func (v Validator) DefaultFloat(f *float64, val float64) {
	if f == nil || *f == 0 {
		*f = val
	}
}

// Sets a default value for a pointer to a string.
// if value does not exists, then set the default specified as the new value.
func (v Validator) DefaultString(s *string, val string) {
	if s == nil {
		*s = val
	}
}
