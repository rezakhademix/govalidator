package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MAC(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "test a correct mac address string value will pass validation",
			field:       "mac",
			value:       "00-B0-D0-63-C2-26",
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "test a correct mac address string value will pass validation",
			field:       "mac",
			value:       "00:B0:D0:63:C2:26",
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "test a wrong string value will fail mac address validation",
			field:       "mac_address",
			value:       "22-aha-sh1-13-b12",
			isPassed:    false,
			message:     "",
			expectedMsg: "mac_address is not valid",
		},
		{
			name:        "test a wrong string value will fail mac address validation",
			field:       "mac",
			value:       "12-$13-me-%3$-e23",
			isPassed:    false,
			message:     "",
			expectedMsg: "mac is not valid",
		},
		{
			name:        "test an empty string value will fail mac validation",
			field:       "mac",
			value:       "",
			isPassed:    false,
			message:     "mac is not valid",
			expectedMsg: "mac is not valid",
		},
		{
			name:        "test an empty space string value will fail mac validation",
			field:       "mac_address",
			value:       " ",
			isPassed:    false,
			message:     "mac_address is not valid",
			expectedMsg: "mac_address is not valid",
		},
	}

	for _, test := range tests {
		v := New()

		v.MAC(test.value, test.field, test.message)

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
