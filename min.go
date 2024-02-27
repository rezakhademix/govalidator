package validator

func (v *Validator) MinInt(i, min int, field, msg string) *Validator {
	v.Check(i >= min, field, msg)

	return v
}
