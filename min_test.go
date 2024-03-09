package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_MinInt(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       int
		min         int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test integer value of `2` will pass if the validator minimum acceptable value is 1",
			field:       "number",
			value:       2,
			min:         1,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test integer value of `3` won't pass if the validator minimum acceptable value is 1",
			field:       "score",
			value:       3,
			min:         10,
			isPassed:    false,
			msg:         "",
			expectedMsg: "score should be more than 10",
		},
		{
			name:        "test integer value of `17` won't pass if the validator minimum acceptable value is 18",
			field:       "age",
			value:       17,
			min:         18,
			isPassed:    false,
			msg:         "age must be greater than 18",
			expectedMsg: "age must be greater than 18",
		},
	}

	v := New()

	for _, test := range tests {
		v.MinInt(test.value, test.min, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed: expected %v, got %v",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func TestValidator_MinFloat(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       float64
		min         float64
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test float value of `2.5` will pass if the validator minimum acceptable value is 1.5",
			field:       "score",
			value:       2.5,
			min:         1.5,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test float value of `21` will pass if the validator minimum acceptable value is 11",
			field:       "score",
			value:       21,
			min:         11,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test float value of `9.75` won't pass if the validator minimum acceptable value is 10",
			field:       "goal",
			value:       9.75,
			min:         10,
			isPassed:    false,
			msg:         "",
			expectedMsg: "goal should be more than 10",
		},
		{
			name:        "test float value of `11.6` won't pass if the validator minimum acceptable value is 121.1",
			field:       "number",
			value:       11.6,
			min:         121.1,
			isPassed:    false,
			msg:         "number must be greater than 121.1",
			expectedMsg: "number must be greater than 121.1",
		},
	}

	v := New()

	for _, test := range tests {
		v.MinFloat(test.value, test.min, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed: expected %v, got %v",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
