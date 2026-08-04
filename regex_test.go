package govalidator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RegexMatches(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		value       string
		pattern     string
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "test an empty string value will fail regex validation",
			field:       "code",
			value:       "",
			pattern:     "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$",
			isPassed:    false,
			message:     "code is not valid",
			expectedMsg: "code is not valid",
		},
		{
			name:        "test an empty space string value will fail regex validation",
			field:       "name",
			value:       " ",
			pattern:     "^[0-9]{10}$",
			isPassed:    false,
			message:     "name is not valid",
			expectedMsg: "name is not valid",
		},
		{
			name:        "test a wrong string value will fail regex validation",
			field:       "id",
			value:       "09377475856",
			pattern:     "^[0-9]{10}$",
			isPassed:    false,
			message:     "id is not valid",
			expectedMsg: "id is not valid",
		},
		{
			name:        "test a correct string value will pass validation",
			field:       "id",
			value:       "1160277052",
			pattern:     "^[0-9]{10}$",
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "test a correct email string value will pass validation",
			field:       "email",
			value:       "rezakhdemix@gmail.com",
			pattern:     "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$",
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
	}

	for _, test := range tests {
		v := New()

		v.RegexMatches(test.value, test.pattern, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equal(
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

// regexCacheLen reports how many compiled patterns are currently cached.
func regexCacheLen() int {
	regexCacheMu.RLock()
	defer regexCacheMu.RUnlock()

	return len(regexCache)
}

func Test_RegexMatches_CacheIsBounded(t *testing.T) {
	v := New()
	for i := 0; i < 300; i++ {
		v.RegexMatches("value", fmt.Sprintf("^pattern-%d$", i), "field", "")
	}

	assert.LessOrEqual(t, regexCacheLen(), maxRegexCacheSize, "regex cache must not grow unbounded")

	v2 := New()
	v2.RegexMatches("abc123", "[a-z]+[0-9]+", "input", "")
	assert.True(t, v2.IsPassed())
}
