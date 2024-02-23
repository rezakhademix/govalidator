package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator_RequiredString(t *testing.T) {
	tests := []struct {
		tag         string
		value       string
		message     string
		isPassed    bool
		exceptedMsg string
	}{
		{
			tag:         "t0",
			value:       "test 0",
			message:     "",
			isPassed:    true,
			exceptedMsg: "",
		},
		{
			tag:         "t1",
			value:       "",
			message:     "t1 is required",
			isPassed:    false,
			exceptedMsg: "t1 is required",
		},
		{
			tag:         "t2",
			value:       " ",
			message:     "t2 is required",
			isPassed:    false,
			exceptedMsg: "t2 is required",
		},
	}

	v := New()
	for _, test := range tests {
		v.RequiredString(test.value, test.tag, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(t, test.exceptedMsg, v.Errors()[test.tag])
		}
	}
}
