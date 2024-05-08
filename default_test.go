package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultInt(t *testing.T) {
	v := New()

	tests := []struct {
		name         string
		actual       int
		expected     int
		defaultValue int
		setDefault   func(*int, int) Validator
	}{
		{
			name:         "Test default int of '3' wont be used if value is greater than 0 or nil",
			actual:       1,
			expected:     1,
			defaultValue: 3,
			setDefault:   v.DefaultInt,
		},
		{
			name:         "Test default int of '3' will be used because value is empty",
			expected:     3,
			defaultValue: 3,
			setDefault:   v.DefaultInt,
		},
	}

	for _, test := range tests {
		test.setDefault(&test.actual, test.defaultValue)

		assert.Equalf(
			t,
			test.actual,
			test.expected,
			"test case %q failed, expected: %s, got: %s",
			test.expected,
			test.actual,
		)
	}
}

func Test_DefaultFloat(t *testing.T) {
	v := New()

	tests := []struct {
		name         string
		actual       float64
		expected     float64
		defaultValue float64
		setDefault   func(*float64, float64) Validator
	}{
		{
			name:         "Test default float of '5.0' won't be used if value is greater than 0.0 or nil",
			actual:       1,
			expected:     1,
			defaultValue: 5.0,
			setDefault:   v.DefaultFloat,
		},
		{
			name:         "Test default float of '1' will be used",
			expected:     1,
			defaultValue: 1,
			setDefault:   v.DefaultFloat,
		},
	}

	for _, test := range tests {
		test.setDefault(&test.actual, test.defaultValue)

		assert.Equalf(
			t,
			test.actual,
			test.expected,
			"test case %q failed, expected: %s, got: %s",
			test.expected,
			test.actual,
		)
	}
}

func Test_DefaultString(t *testing.T) {
	v := New()

	tests := []struct {
		name         string
		actual       string
		expected     string
		defaultValue string
		setDefault   func(*string, string) Validator
	}{
		{
			name:         "Test default string of 'something' won't be used if a value is already valid",
			actual:       "hi",
			expected:     "hi",
			defaultValue: "something",
			setDefault:   v.DefaultString,
		},
		{
			name:         "Test default string of 'hello' will be used",
			expected:     "hello",
			defaultValue: "hello",
			setDefault:   v.DefaultString,
		},
	}

	for _, test := range tests {
		test.setDefault(&test.actual, test.defaultValue)

		assert.Equalf(
			t,
			test.actual,
			test.expected,
			"test case %q failed, expected: %s, got: %s",
			test.expected,
			test.actual,
		)
	}
}
