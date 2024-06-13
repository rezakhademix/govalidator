package govalidator

import "encoding/json"

const (
	// JSON represents rule name which will be used to find the default error message.
	JSON = "json"
	// JSONMsg is the default error message format for fields with JSON validation rule.
	JSONMsg = "%s should be a valid JSON"
)

// IsJSON checks if given string is a valid JSON.
//
// Example:
//
//	v := validator.New()
//	v.IsJSON("{"menu": {"id": "1", "value": "file"}}", "input", "input should be a valid JSON.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) IsJSON(j, field, msg string) Validator {
	v.check(json.Valid([]byte(j)), field, v.msg(JSON, msg, field))

	return v
}
