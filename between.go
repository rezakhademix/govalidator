package validator

const (
	// Between represents rule name which will be used to find the default error message.
	Between = "between"
	// BetweenMsg is the default error message format for fields with Between rule.
	BetweenMsg = "%s should be greater than or equal %v and less than or equal %v"
)

// BetweenInt checks the value under validation to have an integer value between the given min and max.
//
// Example:
//
//	validator.BetweenInt(21, 1, 10, "age", "age must be between 1 and 10.")
func (v *Validator) BetweenInt(i, min, max int, field, msg string) *Validator {
	v.Check(i >= min && i <= max, field, v.msg(Between, msg, field, min, max))

	return v
}

// BetweenFloat checks the field under validation to have a float value between the given min and max.
//
// Example:
//
//	validator.BetweenFloat(3.5, 2.0, 5.0, "height", "height must be between 2.0 and 5.0 meters.")
func (v *Validator) BetweenFloat(f, min, max float64, field, msg string) *Validator {
	v.Check(f >= min && f <= max, field, v.msg(Between, msg, field, min, max))

	return v
}
