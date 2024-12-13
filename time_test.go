package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Time(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		layout      string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test `23:59` will pass under `15:04` layout",
			field:       "arrived_time",
			value:       "23:59",
			layout:      "15:04",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test `11:59PM` will pass under `03:04PM` layout",
			field:       "created_at",
			value:       "11:59PM",
			layout:      "03:04PM",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test `2359` won't pass because it didn't comply defined time layout",
			field:       "due_at",
			value:       "2359",
			layout:      "15:04",
			isPassed:    false,
			msg:         "",
			expectedMsg: "due_at has wrong time format",
		},
	}

	for _, test := range tests {
		v := New()

		v.Time(test.layout, test.value, test.field, test.msg)

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
