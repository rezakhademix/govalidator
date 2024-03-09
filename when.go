package validator

// When method will execute the given closure if the given condition is true.
//
// Example:
//
//	validator.When(len(username) > 0, func() {
//	    validator.RequiredString(username, "username", "username is required.")
//	})
func (v *Validator) When(condition bool, f func()) *Validator {
	if condition {
		f()
	}

	return v
}
