package validator

// MinInt checks i to be greater than min value
func (v *Validator) MinInt(i, min int, field, msg string) *Validator {
	v.Check(i >= min, field, msg)

	return v
}
