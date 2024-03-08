package validator

// When method will execute the given closure if the given condition is true.
func (v *Validator) When(condition bool, f func()) *Validator {
	if condition {
		f()
	}

	return v
}
