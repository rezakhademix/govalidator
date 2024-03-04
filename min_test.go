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
			message:     "",
			isPassed:    false,
			expectedMsg: "t1 should more than 0",
		},
		{
			field:       "t2",
			value:       12,
			min:         20,
			message:     "t2 must be greater than 20",
			isPassed:    false,
			expectedMsg: "t2 must be greater than 20",
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

func TestValidator_MinFloat64(t *testing.T) {
	tests := []struct {
		field       string
		value       float64
		min         float64
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       2.5,
			min:         1.5,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       -0.75,
			min:         -0.25,
			message:     "",
			isPassed:    false,
			expectedMsg: "t1 should more than -0.25",
		},
		{
			field:       "t2",
			value:       1.6,
			min:         7.1,
			message:     "t2 must be greater than 1.6",
			isPassed:    false,
			expectedMsg: "t2 must be greater than 1.6",
		},
	}

	v := New()

	for _, test := range tests {
		v.MinFloat(test.value, test.min, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}
