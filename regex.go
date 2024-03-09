package validator

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
//	validator.RegexMatches("example123", "[a-z]+[0-9]+", "input", "input must contain letters followed by numbers.")
func (v *Validator) RegexMatches(s string, pattern string, field, msg string) *Validator {
	r := regexp.MustCompile(pattern)

	v.Check(r.Match([]byte(s)), field, v.msg(Regex, msg))

	return v
}
