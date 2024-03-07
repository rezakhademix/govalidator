package validator

import "regexp"

const (
	// Regex represents the rule name which will be used to find the default error message.
	Regex = "regex"
	// RegexMsg is the default error message format for fields with the regex validation rule.
	RegexMsg = "%s is not valid"
)

// RegexMatches checks the value of s under validation must match the given regular expression.
func (v *Validator) RegexMatches(s string, pattern string, field, msg string) *Validator {
	r := regexp.MustCompile(pattern)

	v.Check(r.Match([]byte(s)), field, v.msg(Regex, msg))

	return v
}
