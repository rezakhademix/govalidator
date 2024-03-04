package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_MaxInt(t *testing.T) {
	tests := []struct {
		field       string
		value       int
		max         int
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       10,
			max:         10,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       1,
			max:         0,
			message:     "",
			isPassed:    false,
			expectedMsg: "t1 should less than 0",
		},
		{
			field:       "t2",
			value:       122,
			max:         20,
			message:     "t2 must be less than 20",
			isPassed:    false,
			expectedMsg: "t2 must be less than 20",
		},
	}

	v := New()

	for _, test := range tests {
		v.MaxInt(test.value, test.max, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}

func TestValidator_MaxFloat64(t *testing.T) {
	tests := []struct {
		field       string
		value       float64
		max         float64
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       10.1,
			max:         10.8,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       0.1,
			max:         0.01,
			message:     "",
			isPassed:    false,
			expectedMsg: "t1 should less than 0.01",
		},
		{
			field:       "t2",
			value:       122,
			max:         20,
			message:     "t2 must be less than 20",
			isPassed:    false,
			expectedMsg: "t2 must be less than 20",
		},
	}

	v := New()

	for _, test := range tests {
		v.MaxFloat(test.value, test.max, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}
