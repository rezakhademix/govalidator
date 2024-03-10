package govalidator

// When method will execute the given closure if the given condition is true.
//
// Example:
//
//	govalidator.When(len(username) > 0, func() {
//	    govalidator.RequiredString(username, "username", "username is required.")
//	})
func (v *Validator) When(condition bool, f func()) *Validator {
	if condition {
		f()
	}

	return v
}
