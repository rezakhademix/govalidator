package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NotExists(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       any
		table       string
		column      string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test username of joe doesn't exist in defined users table",
			field:       "username",
			value:       "joe",
			table:       "users",
			column:      "username",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test username of reza won't pass validation because it exists in defined users table",
			field:       "username",
			value:       "reza",
			table:       "users",
			column:      "username",
			isPassed:    false,
			msg:         "",
			expectedMsg: "username already exists",
		},
		{
			name:        "test username of reza won't pass validation because it exists in defined users table",
			field:       "username",
			value:       "reza",
			table:       "users",
			column:      "username",
			isPassed:    false,
			msg:         "username `reza` already exists",
			expectedMsg: "username `reza` already exists",
		},
	}

	v := New().
		WithRepo(repo{})

	for _, test := range tests {
		v.NotExists(test.value, test.table, test.column, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if v.IsFailed() {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed: expected: %s, got: %s",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
