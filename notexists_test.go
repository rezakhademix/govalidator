package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidator_NotExists(t *testing.T) {
	tests := []struct {
		tag         string
		value       any
		table       string
		column      string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			tag:         "username",
			value:       "test_2",
			table:       "users",
			column:      "username",
			isPassed:    false,
			msg:         "user with this username already exists",
			expectedMsg: "user with this username already exists",
		},
		{
			tag:         "username",
			value:       "test_3",
			table:       "users",
			column:      "username",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			tag:         "username_2",
			value:       "test_1",
			table:       "users",
			column:      "username",
			isPassed:    false,
			msg:         "",
			expectedMsg: "username_2 already exists",
		},
	}

	v := New().
		WithRepo(repo{})

	for _, test := range tests {
		v.NotExists(test.value, test.table, test.column, test.tag, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if v.IsFailed() {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.tag])
		}
	}
}
