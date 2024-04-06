package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IP4(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "test 192.168.0.1 is a valid ipv4 address",
			field:       "ip",
			value:       "192.168.0.1",
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "test empty string won't pass the ip4 validation rule",
			field:       "ip",
			value:       "",
			isPassed:    false,
			message:     "",
			expectedMsg: "ip should be a valid ipv4",
		},
		{
			name:        "test empty space string won't pass the ip4 validation rule",
			field:       "server_ip",
			value:       "",
			isPassed:    false,
			message:     "",
			expectedMsg: "server_ip should be a valid ipv4",
		},
		{
			name:        "test 192.168.0.256 won't pass the ip4 validation rule",
			field:       "server_ip",
			value:       "192.168.0.256",
			isPassed:    false,
			message:     "",
			expectedMsg: "server_ip should be a valid ipv4",
		},
		{
			name:        "test 255.255.-1.255 won't pass the ip4 validation rule",
			field:       "server_ip",
			value:       "255.255.-1.255",
			isPassed:    false,
			message:     "server_ip should be a valid ipv4",
			expectedMsg: "server_ip should be a valid ipv4",
		},
	}

	for _, test := range tests {
		v := New()

		v.IP4(test.value, test.field, test.message)

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
