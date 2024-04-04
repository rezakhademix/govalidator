package govalidator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Before(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		beforeValue string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test time instant 2009-11-02 is before than 2012-01-01",
			field:       "birth_date",
			value:       "2009-11-02",
			beforeValue: "2012-01-01",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test time instant 2019-01-01 is after than 2010-01-01 and won't pass validation rule",
			field:       "birth_date",
			value:       "2019-01-01",
			beforeValue: "2010-01-01",
			isPassed:    false,
			msg:         "",
			expectedMsg: "birth_date should be before 2010-01-01 00:00:00 +0000 UTC",
		},
		{
			name:        "test time instant 2022-02-01 is after than 2022-01-01 and won't pass validation rule",
			field:       "birth_date",
			value:       "2022-02-01",
			beforeValue: "2022-01-01",
			isPassed:    false,
			msg:         "birth_date can't be before 2022-01-01.",
			expectedMsg: "birth_date can't be before 2022-01-01.",
		},
	}

	v := New()

	for _, test := range tests {
		value, _ := time.Parse("2006-01-02", test.value)
		beforeValue, _ := time.Parse("2006-01-02", test.beforeValue)

		v.Before(value, beforeValue, test.field, test.msg)

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
