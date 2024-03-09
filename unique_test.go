package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		name           string
		values         []any
		expectedResult bool
	}{
		{
			name:           "test given values are unqiue",
			values:         []any{1, 2, 3, 4, 5},
			expectedResult: true,
		},
		{
			name:           "test give values are not unique",
			values:         []any{1, 2, 3, 4, 4},
			expectedResult: false,
		},
		{
			name:           "test give values are not unique",
			values:         []any{"golang", "golang", "php"},
			expectedResult: false,
		},
		{
			name:           "test give values are not unique",
			values:         []any{" ", " "},
			expectedResult: false,
		},
		{
			name:           "test empty slice has no duplicate value",
			values:         []any{},
			expectedResult: true,
		},
		{
			name:           "test single element slice has no duplicate value",
			values:         []any{1},
			expectedResult: true,
		},
		{
			name:           "test single element slice has no duplicate value",
			values:         []any{""},
			expectedResult: true,
		},
	}

	for _, test := range tests {
		result := Unique(test.values)

		assert.Equalf(
			t,
			test.expectedResult,
			result,
			"test case %q failed: expected: %s, got: %s",
			test.name, test.expectedResult, result,
		)
	}
}
