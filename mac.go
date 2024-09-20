package govalidator

const (
	// MAC represents rule name which will be used to find the default error message.
	MAC = "MAC"
	// MACMsg is the default error message format for fields with MAC validation rule.
	MACMsg = "%s is not valid"
	// MACRegex is the default pattern to validate MAC address field.
	MACRegex = "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"
)

// MAC checks the value under validation must match the MACRegex regular expression.
//
// Example:
//
//	v := validator.New()
//	v.MAC("01:23:45:67:89:A", "MAC", "MAC address is not valid.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) MAC(s, field, msg string) Validator {
	v.RegexMatches(s, MACRegex, field, v.msg(MAC, msg, field))

	return v
}
