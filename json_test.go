package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsJSON(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "test {\"menu\": {\"id\": \"1\", \"value\": \"file\"}} is a valid JSON input",
			field:       "input",
			value:       "{\"menu\": {\"id\": \"1\", \"value\": \"file\"}}",
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "test empty string won't pass the JSON validation rule",
			field:       "input",
			value:       "",
			isPassed:    false,
			message:     "",
			expectedMsg: "input should be a valid JSON",
		},
		{
			name:        "test empty space string won't pass the JSON validation rule",
			field:       "input",
			value:       " ",
			isPassed:    false,
			message:     "",
			expectedMsg: "input should be a valid JSON",
		},
		{
			name:        "test `Reza` won't pass the JSON validation rule",
			field:       "input",
			value:       "Reza",
			isPassed:    false,
			message:     "",
			expectedMsg: "input should be a valid JSON",
		},
		{
			name:        "test `{\"example\":2:]}}` won't pass the JSON validation rule",
			field:       "input",
			value:       "{\"example\":2:]}}",
			isPassed:    false,
			message:     "input should be a valid JSON",
			expectedMsg: "input should be a valid JSON",
		},
	}

	for _, test := range tests {
		v := New()

		v.IsJSON(test.value, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed(), test.name)

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
