package validator

const (
	// Exists represents the rule name which will be used to find the default error message.
	Exists = "exists"
	// ExistsMsg is default error message format for records that does not exist.
	ExistsMsg = "%s does not exist"
)

// Exists checks if given value exists in the desired table or not.
//
// Example:
//
//	validator.Exists(42, "users", "id", "user_id", "user with id 42 does not exist.")
func (v *Validator) Exists(value any, table, column, field, msg string) *Validator {
	v.Check(v.repo.Exists(value, table, column), field, v.msg(Exists, msg, field))

	return v
}
