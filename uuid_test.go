package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_UUID(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test a valid uuid will pass validation",
			field:       "uuid",
			value:       "72803040-133f-4295-b85e-c4a98fb8bcad",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test a fake uuid won't pass validation",
			field:       "uuid",
			value:       "98759454-4-504-4295-8543-98763689392",
			isPassed:    false,
			msg:         "uuid is not a valid UUID",
			expectedMsg: "uuid is not a valid UUID",
		},
		{
			name:        "test an empty string won't pass validation",
			field:       "uuid",
			value:       "",
			isPassed:    false,
			msg:         "",
			expectedMsg: "uuid is not a valid UUID",
		},
		{
			name:        "test an empty space string won't pass validation",
			field:       "uuid",
			value:       " ",
			isPassed:    false,
			msg:         "",
			expectedMsg: "uuid is not a valid UUID",
		},
		{
			name:        "test an invalid uuid won't pass validation",
			field:       "uuid",
			value:       "f74c2a2284b3-aa59-504-4295-8543-f74c2a2284b3",
			isPassed:    false,
			msg:         "uuid is not a valid UUID",
			expectedMsg: "uuid is not a valid UUID",
		},
	}

	v := New()

	for _, test := range tests {
		v.UUID(test.value, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expectedMsg: %s, msg: %s",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}

}
