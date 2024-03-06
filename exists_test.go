package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tables = map[string][]map[string]any{
		"users": {
			{
				"id":       1,
				"username": "test_1",
			},
			{
				"id":       2,
				"username": "test_2",
			},
		},
	}
)

type repo struct{}

func (repo) Exists(value any, table, column string) bool {
	data, exists := tables[table]
	if !exists {
		return false
	}

	for _, item := range data {
		if item[column] == value {
			return true
		}
	}

	return false
}

func TestValidator_Exists(t *testing.T) {
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
			tag:         "id",
			value:       5,
			table:       "users",
			column:      "id",
			isPassed:    false,
			msg:         "record not exists",
			expectedMsg: "record not exists",
		},
		{
			tag:         "username",
			value:       "test_1",
			table:       "users",
			column:      "username",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			tag:         "username_2",
			value:       "test_5",
			table:       "users",
			column:      "username",
			isPassed:    false,
			msg:         "",
			expectedMsg: "username_2 not exists",
		},
	}

	v := New().
		WithRepo(repo{})

	for _, test := range tests {
		v.Exists(test.value, test.table, test.column, test.tag, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if v.IsFailed() {
			assert.Equal(t, test.expectedMsg, v.Errors()[test.tag])
		}
	}
}
