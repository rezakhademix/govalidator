package govalidator

// CustomRule is a dynamic method to define any custom validation rule by passing a rule as a function or expression
// which will return a boolean.
func (v Validator) CustomRule(ok bool, field, msg string) Validator {
	if !ok {
		v.addError(field, msg)
	}

	return v
}
