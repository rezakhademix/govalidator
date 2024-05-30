package govalidator

import "time"

const (
	// Before represents rule name which will be used to find the default error message.
	Before = "before"
	// BeforeMsg is the default error message format for fields with Before validation rule.
	BeforeMsg = "%s should be before %v"
)

// Before checks if given time instant t is before u.
//
// Example:
//
//	v := validator.New()
//
//	t, _ := time.Parse("2006-01-02", "2009-01-02") // error ignored for simplicity
//	u, _ := time.Parse("2006-01-02", "2012-01-01") // error ignored for simplicity
//
//	v.Before(t, u, "birth_date", "birth_date should be before 2012-01-01.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) Before(t, u time.Time, field, msg string) Validator {
	v.check(t.Before(u), field, v.msg(Before, msg, field, u))

	return v
}
