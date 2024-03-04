package validator

import (
	"fmt"
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

func TestValidator_RequiredFloat(t *testing.T) {
	tests := []struct {
		tag         string
		value       float64
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			tag:         "t0",
			value:       1.0,
			message:     "",
			isPassed:    true,
			expectedMsg: "f2 is required",
		},
		{
			tag:         "t1",
			value:       -1.0,
			message:     "f2 is required",
			isPassed:    true,
			expectedMsg: fmt.Sprintf(RequiredMsg, "t1"),
		},
		{
			tag:         "t2",
			value:       0.0,
			message:     "f2 is required",
			isPassed:    false,
			expectedMsg: "f2 is required",
		},
	}

	v := New()

	for _, test := range tests {
		v.RequiredFloat(test.value, test.tag, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.tag])
		}
	}
}
