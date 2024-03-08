package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_When(t *testing.T) {
	tests := []struct {
		name         string
		condition    bool
		expectedExec bool
	}{
		{
			name:         "test the given closure will run if the condition is true",
			condition:    len("reza") > 0,
			expectedExec: true,
		},
		{
			name:         "test the given closure won't run if the condition is false",
			condition:    len("reza") < 2,
			expectedExec: false,
		},
		{
			name:         "test the given closure won't run if the condition is false",
			condition:    "" == "golang",
			expectedExec: false,
		},
	}

	v := New()

	for _, test := range tests {
		executed := false

		v.When(test.condition, func() {
			executed = true
		})

		assert.Equalf(
			t,
			test.expectedExec,
			executed,
			"assertion failed, expectedMsg: %s, validatorMsg: %s",
			test.expectedExec,
			executed,
		)
	}
}
