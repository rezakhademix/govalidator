package govalidator

import "regexp"

const (
	// Regex represents rule name which will be used to find the default error message.
	Regex = "regex"
	// RegexMsg is the default error message format for fields with Regex validation rule.
	RegexMsg = "%s is not valid"
)

// RegexMatches checks if the given value of s under validation matches the given regular expression pattern.
//
// Example:
//
//	v := validator.New()
//	v.RegexMatches("example123", "[a-z]+[0-9]+", "input", "input must contain letters followed by numbers.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) RegexMatches(s string, pattern string, field, msg string) Validator {
	r := regexp.MustCompile(pattern)

	v.check(r.Match([]byte(s)), field, v.msg(Regex, msg))

	return v
}
