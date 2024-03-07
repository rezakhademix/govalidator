package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_In(t *testing.T) {
	tests := []struct {
		name             string
		value            any
		acceptableValues []any
		expectedResult   bool
	}{
		{
			name:             "test integer value does not exist in acceptableValues",
			value:            4,
			acceptableValues: []any{1, 2, 3},
			expectedResult:   false,
		},
		{
			name:             "test string value does not exist in acceptable values",
			value:            "redis",
			acceptableValues: []any{"mysql", "mariadb", "postgres"},
			expectedResult:   false,
		},
		{
			name:             "test empty string value does not exist in acceptable values",
			value:            "",
			acceptableValues: []any{"pen", "pencil", "pipe"},
			expectedResult:   false,
		},
		{
			name:             "test empty space string value does not exist in acceptable values",
			value:            " ",
			acceptableValues: []any{"joe", "jane", "john"},
			expectedResult:   false,
		},
		{
			name:             "test integer value exists in acceptable values",
			value:            20,
			acceptableValues: []any{10, 20, 30},
			expectedResult:   true,
		},
		{
			name:             "test string value exists in acceptable values",
			value:            "go",
			acceptableValues: []any{"go", "php", "java"},
			expectedResult:   true,
		},
	}

	for _, test := range tests {
		result := In(test.value, test.acceptableValues...)

		assert.Equalf(
			t,
			test.expectedResult,
			result,
			"test case %q failed: expected %v, got %v",
			test.name, test.expectedResult, result,
		)
	}
}
