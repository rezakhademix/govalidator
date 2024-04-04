package govalidator

import (
	"net/url"
)

const (
	// Url represents rule name which will be used to find the default error message.
	Url = "url"
	// UrlMsg is the default error message format for fields with Url validation rule.
	UrlMsg = "%s should be a valid url"
)

// Url checks if a string value is a valid url or not.
//
// Example:
//
//	v := validator.New()
//	v.Url("https://go.dev/play", "path", "path should be a valid url.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v *Validator) Url(s, field, msg string) *Validator {
	u, err := url.Parse(s)

	v.Check(err == nil && u.Scheme != "" && u.Host != "", field, v.msg(Url, msg, field))

	return v
}
