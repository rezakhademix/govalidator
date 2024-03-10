package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BetweenInt(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       int
		min         int
		max         int
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test integer value of `7` is within [-7, 10] range",
			field:       "score",
			value:       7,
			min:         -7,
			max:         10,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test integer value of `0` is within [0, 10] range",
			field:       "number",
			value:       0,
			min:         0,
			max:         10,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test integer value of `15` is not within [18, 70] range",
			field:       "age",
			value:       15,
			min:         18,
			max:         70,
			isPassed:    false,
			msg:         "",
			expectedMsg: "age should be greater than or equal 18 and less than or equal 70",
		},
		{
			name:        "test integer value of `99` is not within [100, 200] range with custom msg",
			field:       "number",
			value:       99,
			min:         100,
			max:         200,
			isPassed:    false,
			msg:         "number should be greater than or equal 100 and less than or equal 200",
			expectedMsg: "number should be greater than or equal 100 and less than or equal 200",
		},
	}

	v := New()

	for _, test := range tests {
		v.BetweenInt(test.value, test.min, test.max, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

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

func Test_BetweenFloat(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       float64
		min         float64
		max         float64
		msg         string
		isPassed    bool
		expectedMsg string
	}{
		{
			name:        "test float value of `5.33` is within [5.33, 5.38] range",
			field:       "goal",
			value:       5.33,
			min:         5.33,
			max:         5.38,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test float value of `0` is within [-3, 6] range",
			field:       "score",
			value:       0,
			min:         -3,
			max:         6,
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test float value of `18.5` is not within [-5, 14.2] range",
			field:       "number",
			value:       18.5,
			min:         -5,
			max:         14.2,
			isPassed:    false,
			msg:         "",
			expectedMsg: "number should be greater than or equal -5 and less than or equal 14.2",
		},
		{
			name:        "test float value of `122.8` is not within [10.5, 12.2] range",
			field:       "score",
			value:       122.8,
			min:         10.5,
			max:         12.2,
			isPassed:    false,
			msg:         "score should be greater than or equal 10.5 and less than or equal 12.2",
			expectedMsg: "score should be greater than or equal 10.5 and less than or equal 12.2",
		},
	}

	v := New()

	for _, test := range tests {
		v.BetweenFloat(test.value, test.min, test.max, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed())

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
