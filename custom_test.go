package govalidator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CustomRule(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       bool
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test CustomRule with true condition",
			value:       true,
			field:       "username",
			msg:         "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			name:        "test CustomRule with false condition and custom message",
			value:       false,
			field:       "username",
			msg:         "username must be unique",
			isPassed:    false,
			expectedMsg: "username must be unique",
		},
		{
			name:        "test CustomRule with false condition and empty message",
			value:       false,
			field:       "email",
			msg:         "",
			isPassed:    false,
			expectedMsg: "",
		},
		{
			name:        "test CustomRule with true condition and custom message",
			value:       true,
			field:       "email",
			msg:         "email must be valid",
			isPassed:    true,
			expectedMsg: "",
		},
	}

	for _, test := range tests {
		v := New()

		v.CustomRule(test.value, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed(), test.name)

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
