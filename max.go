package validator

// MaxInt checks i to be less than max value
func (v *Validator) MaxInt(i, max int, field, msg string) *Validator {
	v.Check(i <= max, field, msg)

	return v
}
