package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_MinInt(t *testing.T) {
	tests := []struct {
		field       string
		value       int
		min         int
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       2,
			min:         1,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       -1,
			min:         0,
			message:     "t1 must be greater than 0",
			isPassed:    false,
			expectedMsg: "t1 must be greater than 0",
		},
		{
			field:       "t2",
			value:       12,
			min:         20,
			message:     "t1 must be greater than 20",
			isPassed:    false,
			expectedMsg: "t1 must be greater than 20",
		},
	}

	v := New()

	for _, test := range tests {
		v.MinInt(test.value, test.min, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}
