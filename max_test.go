package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxInt(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       int
		max         int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test integer value of `10` will pass if the validator defined maximum acceptable value to 10",
			field:       "score",
			value:       10,
			max:         10,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test integer value of `71` won't pass because the validator defined maximum acceptable value to 50",
			field:       "age",
			value:       71,
			max:         50,
			isPassed:    false,
			msg:         "",
			expectedMsg: "age should be less than 50",
		},
		{
			name:        "test integer value of `141` won't pass because the validator defined maximum acceptable value to 20 with custom msg",
			field:       "goal",
			value:       141,
			max:         20,
			isPassed:    false,
			msg:         "goal have to be less than 20",
			expectedMsg: "goal have to be less than 20",
		},
	}

	for _, test := range tests {
		v := New()

		v.MaxInt(test.value, test.max, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
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

func Test_MaxFloat64(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       float64
		max         float64
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test float value of `20.4` will pass if the validator defined maximum acceptable value to 110.3",
			field:       "score",
			value:       20.4,
			max:         110.3,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test float value of `0.1` won't pass because the validator defined maximum acceptable value to 0.01",
			field:       "score",
			value:       0.1,
			max:         0.01,
			isPassed:    false,
			msg:         "",
			expectedMsg: "score should be less than 0.01",
		},
		{
			name:        "test float value of `141` won't pass because the validator defined maximum acceptable value to 20 with custom msg",
			field:       "goal",
			value:       122.23,
			max:         20,
			isPassed:    false,
			msg:         "goal have to be less than 20",
			expectedMsg: "goal have to be less than 20",
		},
	}

	for _, test := range tests {
		v := New()

		v.MaxFloat(test.value, test.max, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
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

func Test_MaxString(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		max         int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test `rey` will pass validation when maximum valid length is 5",
			field:       "name",
			value:       "rey",
			max:         5,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test empty string will pass validation when maximum valid length is 2",
			field:       "username",
			value:       "",
			max:         2,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test empty space string won't pass validation when maximum valid length is 0",
			field:       "username",
			value:       " ",
			max:         -1,
			isPassed:    false,
			msg:         "",
			expectedMsg: "username should have less than -1 characters",
		},
		{
			name:        "test `abcd` won't pass validation when maximum valid length is 3",
			field:       "alphabet",
			value:       "abcd",
			max:         3,
			isPassed:    false,
			msg:         "alphabet should have less than 3 characters",
			expectedMsg: "alphabet should have less than 3 characters",
		},
	}

	for _, test := range tests {
		v := New()

		v.MaxString(test.value, test.max, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
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
