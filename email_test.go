package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_Email(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		pattern     string
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "test an empty string value will fail email validation",
			field:       "email",
			value:       "",
			pattern:     EmailRegex,
			isPassed:    false,
			message:     "email is not valid",
			expectedMsg: "email is not valid",
		},
		{
			name:        "test an empty space string value will fail email validation",
			field:       "email_address",
			value:       " ",
			pattern:     EmailRegex,
			isPassed:    false,
			message:     "email_address is not valid",
			expectedMsg: "email_address is not valid",
		},
		{
			name:        "test a wrong string value will fail email validation",
			field:       "email",
			value:       "09377475856",
			pattern:     EmailRegex,
			isPassed:    false,
			message:     "email is not valid",
			expectedMsg: "email is not valid",
		},
		{
			name:        "test a wrong string value will fail email validation",
			field:       "email",
			value:       "^$*me%$e.com",
			pattern:     EmailRegex,
			isPassed:    false,
			message:     "email is not valid",
			expectedMsg: "email is not valid",
		},
		{
			name:        "test a correct email string value will pass validation",
			field:       "email",
			value:       "rezakhdemix@gmail.com",
			pattern:     EmailRegex,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
	}

	for _, test := range tests {
		v := New()

		v.RegexMatches(test.value, test.pattern, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"assertion failed, expectedMsg: %s, validatorMsg: %s",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
