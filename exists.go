package govalidator

const (
	// Exists represents rule name which will be used to find the default error message.
	Exists = "exists"
	// ExistsMsg is default error message format for fields with Exists validation rule.
	ExistsMsg = "%s does not exist"
)

// Exists checks if given value exists in the desired table or not.
//
// Example:
//
//	v := validator.New()
//	v.Exists(42, "users", "id", "user_id", "user with id 42 does not exist.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) Exists(value any, table, column, field, msg string) Validator {
	v.check(v.repo.Exists(value, table, column), field, v.msg(Exists, msg, field))

	return v
}
