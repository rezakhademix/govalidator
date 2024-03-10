package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LenString(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		len         int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test the length of `abcd` as a string is 4",
			field:       "alphabet",
			value:       "abcd",
			len:         4,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test the length of `xyz` won't pass validation because validation expect 7 characters",
			field:       "alphabet",
			value:       "xyz",
			len:         7,
			isPassed:    false,
			msg:         "",
			expectedMsg: "alphabet should be 7 characters",
		},
		{
			name:        "test the length of ` 2345` won't pass validation because validation expect 5 characters",
			field:       "number",
			value:       " 2345",
			len:         5,
			isPassed:    false,
			msg:         "number should have 5 characters",
			expectedMsg: "number should have 5 characters",
		},
	}

	v := New()

	for _, test := range tests {
		v.LenString(test.value, test.len, test.field, test.msg)

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

func Test_LenInt(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       int
		len         int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test `1234567` will pass as a string with length of 6",
			field:       "numbers",
			value:       123456,
			len:         6,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test `123` won't pass validation as string, because validation expect 7 characters",
			field:       "score",
			value:       123,
			len:         7,
			isPassed:    false,
			msg:         "",
			expectedMsg: "score should be 7 characters",
		},
		{
			name:        "test `1223` won't pass validation as string, because validation expect 10 characters",
			field:       "score",
			value:       1223,
			len:         10,
			isPassed:    false,
			msg:         "score should have 10 characters",
			expectedMsg: "score should have 10 characters",
		},
	}

	v := New()

	for _, test := range tests {
		v.LenInt(test.value, test.len, test.field, test.msg)

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

func Test_LenSlice(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       []any
		len         int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test a slice of any with in values and len 3 will pass defined validation with len 3",
			field:       "ages",
			value:       []any{21, 31, 41},
			len:         3,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test a slice of any with float values and len 4 will pass defined validation with len 4",
			field:       "averages",
			value:       []any{11.1, 32.2, 73.9, 12.23},
			len:         4,
			isPassed:    true,
			msg:         "",
			expectedMsg: "%s should have %d items",
		},
		{
			name:        "test a slice of string values and len 3 won't pass because govalidator wants a slice with len of 11",
			field:       "names",
			value:       []any{"sanchez", "marco", "lito"},
			len:         11,
			isPassed:    false,
			msg:         "names must have len of 11",
			expectedMsg: "names must have len of 11",
		},
	}

	v := New()

	for _, test := range tests {
		v.LenSlice(test.value, test.len, test.field, test.msg)

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
