package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Email(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "test a correct email string value will pass validation",
			field:       "email",
			value:       "rezakhdemix@gmail.com",
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "test a wrong string value will fail email validation",
			field:       "email",
			value:       "09377475856",
			isPassed:    false,
			message:     "",
			expectedMsg: "email is not valid",
		},
		{
			name:        "test a wrong string value will fail email validation",
			field:       "email",
			value:       "^$*me%$e.com",
			isPassed:    false,
			message:     "",
			expectedMsg: "email is not valid",
		},
		{
			name:        "test an empty string value will fail email validation",
			field:       "email",
			value:       "",
			isPassed:    false,
			message:     "email is not valid",
			expectedMsg: "email is not valid",
		},
		{
			name:        "test an empty space string value will fail email validation",
			field:       "email_address",
			value:       " ",
			isPassed:    false,
			message:     "email_address is not valid",
			expectedMsg: "email_address is not valid",
		},
	}

	v := New()

	for _, test := range tests {
		v.Email(test.value, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
