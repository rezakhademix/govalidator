package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultInt(t *testing.T) {
	v := New()

	tests := []struct {
		actual       int
		expected     int
		defaultValue int
		setDefault   func(*int, int)
	}{
		{
			actual:       1,
			expected:     1,
			defaultValue: 3,
			setDefault:   v.DefaultInt,
		},
		{
			expected:     3,
			defaultValue: 3,
			setDefault:   v.DefaultInt,
		},
	}

	for _, test := range tests {
		test.setDefault(&test.actual, test.defaultValue)

		assert.Equal(t, test.actual, test.expected)
	}
}

func Test_DefaultFloat(t *testing.T) {
	v := New()

	tests := []struct {
		actual       float64
		expected     float64
		defaultValue float64
		setDefault   func(*float64, float64)
	}{
		{
			actual:       1,
			expected:     1,
			defaultValue: 5.0,
			setDefault:   v.DefaultFloat,
		},
		{
			expected:     1,
			defaultValue: 1,
			setDefault:   v.DefaultFloat,
		},
	}

	for _, test := range tests {
		test.setDefault(&test.actual, test.defaultValue)

		assert.Equal(t, test.actual, test.expected)
	}
}
