package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator_BetweenInt(t *testing.T) {
	tests := []struct {
		field       string
		value       int
		min         int
		max         int
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       7,
			min:         -7,
			max:         10,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       -21,
			min:         -1,
			max:         -5,
			message:     "",
			isPassed:    false,
			expectedMsg: fmt.Sprintf(BetweenMsg, "t1", -1, -5),
		},
		{
			field:       "t2",
			value:       90,
			min:         0,
			max:         9,
			message:     "t2 must be larger than 0 & less than 9",
			isPassed:    false,
			expectedMsg: "t2 must be larger than 0 & less than 9",
		},
	}

	v := New()

	for _, test := range tests {
		v.BetweenInt(test.value, test.min, test.max, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())
		if v.IsFailed() {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}

func TestValidator_BetweenFloat64(t *testing.T) {
	tests := []struct {
		field       string
		value       float64
		min         float64
		max         float64
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       5.33,
			min:         1.21,
			max:         5.38,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       -5.31,
			min:         -1.5,
			max:         -0.4,
			message:     "",
			isPassed:    false,
			expectedMsg: fmt.Sprintf(BetweenMsg, "t1", -1.5, -0.4),
		},
		{
			field:       "t2",
			value:       90,
			min:         0,
			max:         9,
			message:     "t2 must be larger than 0 & less than 9",
			isPassed:    false,
			expectedMsg: "t2 must be larger than 0 & less than 9",
		},
	}

	v := New()

	for _, test := range tests {
		v.BetweenFloat64(test.value, test.min, test.max, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())
		if v.IsFailed() {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}
