package validator

const (
	// Min represents the rule name which will be used to find the default error message.
	Min = "min"
	// MinMsg is the default error message format for fields with the minimum validation rule.
	MinMsg = "%s should be more than %v"
)

// MinInt checks if the given integer value is greater than or equal the given min value.
//
// Example:
//
//	validator.MinInt(18, 0, "age", "age must be at least 0.")
func (v *Validator) MinInt(i, min int, field, msg string) *Validator {
	v.Check(i >= min, field, v.msg(Min, msg, field, min))

	return v
}

// MinFloat checks if the given float value is greater than or equal the given min value.
//
// Example:
//
//	validator.MinFloat(5.0, 0.0, "height", "height must be at least 0.0 meters.")
func (v *Validator) MinFloat(f, min float64, field, msg string) *Validator {
	v.Check(f >= min, field, v.msg(Min, msg, field, min))

	return v
}
