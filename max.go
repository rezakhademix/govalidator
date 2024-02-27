package validator

func (v *Validator) MaxInt(i, max int, field, msg string) *Validator {
	v.Check(i <= max, field, msg)

	return v
}
