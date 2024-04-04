package govalidator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_After(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		afterValue  string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test time instant 2012-01-01 is after 2009-11-02",
			field:       "birth_date",
			value:       "2012-01-01",
			afterValue:  "2009-11-02",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test time instant 2010-01-01 isn't after 2019-01-01 and won't pass validation rule",
			field:       "birth_date",
			value:       "2010-01-01",
			afterValue:  "2019-01-01",
			isPassed:    false,
			msg:         "",
			expectedMsg: "birth_date should be after 2019-01-01 00:00:00 +0000 UTC",
		},
		{
			name:        "test time instant 2022-01-01 isn't after 2022-02-01 and won't pass validation rule",
			field:       "birth_date",
			value:       "2022-01-01",
			afterValue:  "2022-02-01",
			isPassed:    false,
			msg:         "birth_date can't be after 2022-02-01.",
			expectedMsg: "birth_date can't be after 2022-02-01.",
		},
	}

	v := New()

	for _, test := range tests {
		value, _ := time.Parse("2006-01-02", test.value)
		afterValue, _ := time.Parse("2006-01-02", test.afterValue)

		v.After(value, afterValue, test.field, test.msg)

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
