package validator

const (
	// NotExists represents the rule name which will be used to find the default error message.
	NotExists = "notExists"

	// NotExistsMsg is default error message format for records that don't exist.
	NotExistsMsg = "%s already exists"
)

// NotExists checks if given value doesn't exist in desired table or not.
func (v *Validator) NotExists(value any, table, column, field, msg string) *Validator {
	v.Check(!v.repo.Exists(value, table, column), field, v.msg(NotExists, msg, field))

	return v
}
