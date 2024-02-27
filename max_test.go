package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_MaxInt(t *testing.T) {
	tests := []struct {
		field        string
		value        int
		max          int
		message      string
		isPassed     bool
		excpectedMsg string
	}{
		{
			field:        "t0",
			value:        10,
			max:          10,
			message:      "",
			isPassed:     true,
			excpectedMsg: "",
		},
		{
			field:        "t1",
			value:        1,
			max:          0,
			message:      "t1 must be less than 0",
			isPassed:     false,
			excpectedMsg: "t1 must be less than 0",
		},
		{
			field:        "t2",
			value:        122,
			max:          20,
			message:      "t1 must be less than 20",
			isPassed:     false,
			excpectedMsg: "t1 must be less than 20",
		},
	}

	v := New()

	for _, test := range tests {
		v.MaxInt(test.value, test.max, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.excpectedMsg, v.Errors()[test.field])
		}
	}
}
