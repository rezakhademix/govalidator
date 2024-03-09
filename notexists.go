package validator

const (
	// NotExists represents rule name which will be used to find the default error message.
	NotExists = "notExists"

	// NotExistsMsg is default error message format for fields with with NotExists validation rule.
	NotExistsMsg = "%s already exists"
)

// NotExists checks if the given value doesn't exist in the desired table.
//
// Example:
//
//	validator.NotExists(42, "users", "id", "user_id", "user with id 42 already exists.")
func (v *Validator) NotExists(value any, table, column, field, msg string) *Validator {
	v.Check(!v.repo.Exists(value, table, column), field, v.msg(NotExists, msg, field))

	return v
}
