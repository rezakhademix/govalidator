package govalidator

import "time"

const (
	// After represents rule name which will be used to find the default error message.
	After = "after"
	// AfterMsg is the default error message format for fields with After validation rule.
	AfterMsg = "%s should be after %v"
)

// After checks if given time instant t is after u.
//
// Example:
//
//	v := validator.New()
//
//	t, _ := time.Parse("2006-01-02", "2009-01-02") // error ignored for simplicity
//	u, _ := time.Parse("2006-01-02", "2012-01-01") // error ignored for simplicity
//
//	v.After(t, u, "birth_date", "birth_date should be after 2012-01-01.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) After(t, u time.Time, field, msg string) Validator {
	v.check(t.After(u), field, v.msg(After, msg, field, u))

	return v
}
