package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RequiredInt(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test integer value of `21` will pass the required int validation",
			field:       "age",
			value:       21,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test integer value of `0` won't pass the required int validation",
			field:       "age",
			value:       0,
			isPassed:    false,
			msg:         "",
			expectedMsg: "age is required",
		},
		{
			name:        "test integer value of `0` won't pass the required int validation",
			field:       "age",
			value:       0,
			isPassed:    false,
			msg:         "age is required",
			expectedMsg: "age is required",
		},
	}

	for _, test := range tests {
		v := New()

		v.RequiredInt(test.value, test.field, test.msg)

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

func Test_RequiredString(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test string value of `lion` will pass the required string validation",
			field:       "animal",
			value:       "lion",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test string value of `1160277052` will pass the required string validation",
			field:       "id",
			value:       "1160277052",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test emtpy string value won't pass the required string validation",
			field:       "name",
			value:       "",
			isPassed:    false,
			msg:         "name is required",
			expectedMsg: "name is required",
		},
		{
			name:        "test emtpy space string value won't pass the required string validation",
			field:       "last_name",
			value:       " ",
			isPassed:    false,
			msg:         "last_name is required",
			expectedMsg: "last_name is required",
		},
	}

	for _, test := range tests {
		v := New()

		v.RequiredString(test.value, test.field, test.msg)

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

func Test_RequiredFloat(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       float64
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test float value of `12.23` will pass the required float validation",
			field:       "average",
			value:       12.23,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test float value of `0` won't pass the required float validation",
			field:       "score",
			value:       0,
			isPassed:    false,
			msg:         "",
			expectedMsg: "score is required",
		},
		{
			name:        "test float value of `0` won't pass the required float validation",
			field:       "number",
			value:       0,
			isPassed:    false,
			msg:         "number must be passed",
			expectedMsg: "number must be passed",
		},
	}

	for _, test := range tests {
		v := New()

		v.RequiredFloat(test.value, test.field, test.msg)

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

func Test_RequiredSlice(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       []any
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test slice of `[]{1, 2, 3}` will pass the required slice validation",
			field:       "scores",
			value:       []any{1, 2, 3},
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test slice of `{2.2, 12.6, 13.6}` will pass the required slice validation",
			field:       "averages",
			value:       []any{2.2, 12.6, 13.6},
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test emtpy slice won't pass the required slice validation",
			field:       "scores",
			value:       []any{},
			isPassed:    false,
			msg:         "",
			expectedMsg: "scores is required",
		},
		{
			name:        "test slice of names won't pass the required slice validation",
			field:       "names",
			value:       []any{},
			isPassed:    false,
			msg:         "names is required",
			expectedMsg: "names is required",
		},
	}

	for _, test := range tests {
		v := New()

		v.RequiredSlice(test.value, test.field, test.msg)

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
