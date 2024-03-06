package validator

const (
	// NotExists represents the rule name which will be used to find the default error message.
	NotExists = "notExists"

	// NotExistsMsg is default error message format for records that already exist.
	NotExistsMsg = "%s already exists"
)

// NotExists checks value not exists in database.
func (v *Validator) NotExists(value any, table, column, field, msg string) *Validator {
	v.Check(!v.repo.Exists(value, table, column), field, v.msg(NotExists, msg, field))

	return v
}
