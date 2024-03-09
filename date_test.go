package validator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Date(t *testing.T) {
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
			name:        "test `1995-02-12` will pass under `2006-01-02` layout",
			field:       "birth_date",
			value:       "1995-02-12",
			layout:      "2006-01-02",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test `1995/02/12` will pass under 2006/01/02 layout",
			field:       "created_at",
			value:       "1995/02/12",
			layout:      "2006/01/02",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test `15:04:05` will pass under TimeOnly layout",
			field:       "arrived_at",
			value:       "15:04:05",
			layout:      time.TimeOnly,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test `20240207` won't pass because it didn't comply defined date layout",
			field:       "due_at",
			value:       "20240207",
			layout:      "2006-01-02",
			isPassed:    false,
			msg:         "",
			expectedMsg: "due_at has wrong date format",
		},
		{
			name:        "test `1995-02-12` won't pass because it didn't comply defined date layout",
			field:       "expired_at",
			value:       "1995-02-12",
			layout:      "2006/01/02",
			isPassed:    false,
			msg:         "expired_at has wrong date format",
			expectedMsg: "expired_at has wrong date format",
		},
	}

	v := New()

	for _, test := range tests {
		v.Date(test.value, test.layout, test.field, test.msg)

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
