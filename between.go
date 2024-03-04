package validator

const (
	// Between represents the rule name which will be used to find the default error message.
	Between = "between"
	// BetweenMsg is the default error message format for len rule.
	BetweenMsg = "%s should more than %v and less than %v"
)

// BetweenInt checks i to be less than max and more than min value
func (v *Validator) BetweenInt(i, min, max int, field, msg string) *Validator {
	v.Check(i >= min && i <= max, field, v.msg(Between, msg, field, min, max))

	return v
}

// BetweenFloat64 checks i to be less than max and more than min value
func (v *Validator) BetweenFloat64(i, min, max float64, field, msg string) *Validator {
	v.Check(i > min && i < max, field, v.msg(Between, msg, field, min, max))

	return v
}
