package validator

const (
	// Max represents the rule name which will be used to find the default error message.
	Max = "max"
	// MaxMsg is the default error message format for len rule.
	MaxMsg = "%s should less than %v"
)

// MaxInt checks i to be less than max value
func (v *Validator) MaxInt(i, max int, field, msg string) *Validator {
	v.Check(i <= max, field, v.msg(Max, msg, field, max))

	return v
}

// MaxFloat64 checks i to be less than max value
func (v *Validator) MaxFloat64(i, max float64, field, msg string) *Validator {
	v.Check(i <= max, field, v.msg(Max, msg, field, max))

	return v
}
