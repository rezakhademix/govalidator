package govalidator

// When method will execute the given closure if the given condition is true.
//
// Example:
//
//	v := validator.New()
//	v.When(len(username) > 0, func() {
//	    validator.RequiredString(username, "username", "username is required.")
//	})
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) When(condition bool, f func()) Validator {
	if condition {
		f()
	}

	return v
}
