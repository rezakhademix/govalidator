package validator

const (
	// Min represents the rule name which will be used to find the default error message.
	Min = "min"
	// MinMsg is the default error message format for required fields.
	MinMsg = "%s should more than %v"
)

// MinInt checks i to be greater than min value
func (v *Validator) MinInt(i, min int, field, msg string) *Validator {
	v.Check(i >= min, field, v.msg(Min, msg, field, min))

	return v
}

// MinFloat64 checks i to be greater than min value
func (v *Validator) MinFloat64(i, min float64, field, msg string) *Validator {
	v.Check(i >= min, field, v.msg(Min, msg, field, min))

	return v
}
