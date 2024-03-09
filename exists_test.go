package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tables = map[string][]map[string]any{
	"users": {
		{
			"id":       1,
			"username": "reza",
			"nickname": "khademi",
		},
		{
			"id":       2,
			"username": "adel",
			"nickname": "haddadi",
		},
	},
}

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
			name:        "test username of adel exists in defined users table",
			field:       "username",
			value:       "adel",
			table:       "users",
			column:      "username",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test nickname of Horizon does not exist in defined users table",
			field:       "nickname",
			value:       "Horizon",
			table:       "users",
			column:      "nickname",
			isPassed:    false,
			msg:         "",
			expectedMsg: "nickname does not exist",
		},
		{
			name:        "test id of 500 does not exist in defined users table",
			field:       "id",
			value:       500,
			table:       "users",
			column:      "id",
			isPassed:    false,
			msg:         "user with id of 5 does not exist in users table",
			expectedMsg: "user with id of 5 does not exist in users table",
		},
	}

	v := New().
		WithRepo(repo{})

	for _, test := range tests {
		v.Exists(test.value, test.table, test.column, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if v.IsFailed() {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected %v, got %v",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
