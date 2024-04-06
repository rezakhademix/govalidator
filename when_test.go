package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_When(t *testing.T) {
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

	for _, test := range tests {
		v := New()

		executed := false

		v.When(test.condition, func() {
			executed = true
		})

		assert.Equalf(
			t,
			test.expectedExec,
			executed,
			"test case %q failed, expected %s, got: %s",
			test.expectedExec,
			executed,
		)
	}
}
