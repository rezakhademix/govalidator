package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumericString(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test `1234` will pass validation as a numeric-string value",
			field:       "postal_code",
			value:       "1234",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test `abc` won't pass numeric-string validation",
			field:       "postal_code",
			value:       "abc",
			isPassed:    false,
			msg:         "",
			expectedMsg: "postal_code must be a numeric string",
		},
		{
			name:        "test ` ` won't pass numeric-string validation",
			field:       "postal_code",
			value:       " ",
			isPassed:    false,
			msg:         "postal_code must be a numeric string",
			expectedMsg: "postal_code must be a numeric string",
		},
	}

	for _, test := range tests {
		v := New()

		v.NumericString(test.value, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"%q failed: expected: %s, got: %s",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
