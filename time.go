package govalidator

import (
	"time"
)

const (
	// Time represents rule name which will be used to find the default error message.
	Time = "time"
	// TimeMsg is the default error message format for fields with Time validation rule.
	TimeMsg = "%s has wrong time format"
)

// Time checks the value under validation to be a valid, non-relative time with give layout.
//
// Example:
//
//	v := validator.New()
//	v.Time("15:04", "15:00","arrived_time", "arrived_time must be a valid time in the format HH:MM.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) Time(layout, d, field, msg string) Validator {
	_, err := time.Parse(layout, d)
	if err != nil {
		v.check(false, field, v.msg(Time, msg, field))
	}

	return v
}
