package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_RequiredString(t *testing.T) {
	tests := []struct {
		tag         string
		value       string
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			tag:         "t0",
			value:       "test 0",
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			tag:         "t1",
			value:       "",
			message:     "t1 is required",
			isPassed:    false,
			expectedMsg: "t1 is required",
		},
		{
			tag:         "t2",
			value:       " ",
			message:     "t2 is required",
			isPassed:    false,
			expectedMsg: "t2 is required",
		},
	}

	v := New()
	for _, test := range tests {
		v.RequiredString(test.value, test.tag, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.tag])
		}
	}
}
