package govalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Url(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		isPassed    bool
		msg         string
		expectedMsg string
	}{
		{
			name:        "test correct string url will pass the url validation rule",
			field:       "path",
			value:       "https://go.dev",
			isPassed:    true,
			msg:         "",
			expectedMsg: "",
		},
		{
			name:        "test string value of `http://` won't pass the url validation rule",
			field:       "path",
			value:       "http://",
			isPassed:    false,
			msg:         "",
			expectedMsg: "path should be a valid url",
		},
		{
			name:        "test empty string value of `` won't pass the url validation rule",
			field:       "path",
			value:       "",
			isPassed:    false,
			msg:         "",
			expectedMsg: "path should be a valid url",
		},
		{
			name:        "test random string value won't pass the url validation rule",
			field:       "path",
			value:       "qazbvnfhegrtyuomln",
			isPassed:    false,
			msg:         "",
			expectedMsg: "path should be a valid url",
		},
		{
			name:        "test emtpy string value won't pass the url validation rule",
			field:       "url",
			value:       "/foo/bar",
			isPassed:    false,
			msg:         "",
			expectedMsg: "url should be a valid url",
		},
		{
			name:        "test emtpy space string value won't pass the url validation rule",
			field:       "path",
			value:       " ",
			isPassed:    false,
			msg:         "",
			expectedMsg: "path should be a valid url",
		},
		{
			name:        "test a relative url value like: `foo/bar` won't pass the url validation rule",
			field:       "avatar_path",
			value:       "/foo/bar",
			isPassed:    false,
			msg:         "avatar_path must be a valid url",
			expectedMsg: "avatar_path must be a valid url",
		},
	}

	v := New()

	for _, test := range tests {
		v.URL(test.value, test.field, test.msg)

		assert.Equal(t, test.isPassed, v.IsPassed(), test.name)

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed: expected: %s, got: %s",
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}
