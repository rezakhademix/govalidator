package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_msg(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		msg         string
		expectedMsg string
	}{
		{
			name:        "test not exists method will result a panic",
			method:      "qwert",
			msg:         "",
			expectedMsg: "method default validation message does not exist in methodToErrorMessage",
		},
		{
			name:        "test empty string method will result a panic",
			method:      "",
			msg:         "",
			expectedMsg: "method default validation message does not exist in methodToErrorMessage",
		},
		{
			name:        "test empty space string method will result a panic",
			method:      " ",
			msg:         "",
			expectedMsg: "method default validation message does not exist in methodToErrorMessage",
		},
	}

	for _, test := range tests {
		v := New()

		assert.PanicsWithError(t, test.expectedMsg, func() { v.msg(test.method, test.msg) })
	}
}
