package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator_LenString(t *testing.T) {
	tests := []struct {
		field       string
		value       string
		len         int
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       "asdfghj",
			len:         7,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       "asd",
			len:         5,
			message:     "",
			isPassed:    false,
			expectedMsg: fmt.Sprintf(LenMsg, "t1", 5),
		},
		{
			field:       "t2",
			value:       " 2345",
			len:         5,
			message:     "t2 should be 5 character",
			isPassed:    false,
			expectedMsg: "t2 should be 5 character",
		},
	}

	v := New()

	for _, test := range tests {
		v.LenString(test.value, test.len, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())
		if v.IsFailed() {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}

func TestValidator_LenInt(t *testing.T) {
	tests := []struct {
		field       string
		value       int
		len         int
		message     string
		isPassed    bool
		expectedMsg string
	}{
		{
			field:       "t0",
			value:       1234567,
			len:         7,
			message:     "",
			isPassed:    true,
			expectedMsg: "",
		},
		{
			field:       "t1",
			value:       567,
			len:         5,
			message:     "",
			isPassed:    false,
			expectedMsg: fmt.Sprintf(LenMsg, "t1", 5),
		},
	}

	v := New()

	for _, test := range tests {
		v.LenInt(test.value, test.len, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())
		if v.IsFailed() {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.field])
		}
	}
}
