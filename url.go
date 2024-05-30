package govalidator

import (
	"net/url"
)

const (
	// URL represents rule name which will be used to find the default error message.
	URL = "url"
	// URLMsg is the default error message format for fields with Url validation rule.
	URLMsg = "%s should be a valid url"
)

// URL checks if a string value is a valid url or not.
//
// Example:
//
//	v := validator.New()
//	v.URL("https://go.dev/play", "path", "path should be a valid url.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) URL(s, field, msg string) Validator {
	u, err := url.Parse(s)

	v.check(err == nil && u.Scheme != "" && u.Host != "", field, v.msg(URL, msg, field))

	return v
}
