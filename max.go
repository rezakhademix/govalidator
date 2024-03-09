package validator

const (
	// Max represents rule name which will be used to find the default error message.
	Max = "max"
	// MaxMsg is the default error message format for fields with the Max validation rule.
	MaxMsg = "%s should be less than %v"
)

// MaxInt checks if the integer value is less than or equal the given max value.
//
// Example:
//
//	validator.MaxInt(10, 100, "age", "age must be less than 100.")
func (v *Validator) MaxInt(i, max int, field, msg string) *Validator {
	v.Check(i <= max, field, v.msg(Max, msg, field, max))

	return v
}

// MaxFloat checks if the given float value is less than or equal the given max value.
//
// Example:
//
//	validator.MaxFloat(3.5, 5.0, "height", "height must be less than 5.0 meters.")
func (v *Validator) MaxFloat(f, max float64, field, msg string) *Validator {
	v.Check(f <= max, field, v.msg(Max, msg, field, max))

	return v
}
