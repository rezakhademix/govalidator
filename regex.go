package govalidator

import (
	"regexp"
	"sync"
)

const (
	// Regex represents rule name which will be used to find the default error message.
	Regex = "regex"
	// RegexMsg is the default error message format for fields with Regex validation rule.
	RegexMsg = "%s is not valid"
)

// RegexMatches checks if the given value of s under validation matches the given regular expression pattern.
//
// Example:
//
//	v := validator.New()
//	v.RegexMatches("example123", "[a-z]+[0-9]+", "input", "input must contain letters followed by numbers.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) RegexMatches(s string, pattern string, field, msg string) Validator {
	if !compiledRegex(pattern).MatchString(s) {
		v.addError(field, v.msg(Regex, msg))
	}

	return v
}

// maxRegexCacheSize caps how many compiled patterns regexCache may hold, so
// dynamically-built patterns cannot grow the process memory without bound.
const maxRegexCacheSize = 128

// regexCache caches compiled patterns so repeated RegexMatches calls do not
// recompile them. Once maxRegexCacheSize distinct patterns are cached, new
// patterns are still compiled and used but no longer stored.
var (
	regexCacheMu sync.RWMutex
	regexCache   = make(map[string]*regexp.Regexp)
)

// compiledRegex returns the cached compiled form of pattern, compiling it on
// first use. Like regexp.MustCompile, it panics on an invalid pattern.
func compiledRegex(pattern string) *regexp.Regexp {
	regexCacheMu.RLock()
	re, ok := regexCache[pattern]
	regexCacheMu.RUnlock()

	if ok {
		return re
	}

	re = regexp.MustCompile(pattern)

	regexCacheMu.Lock()
	if len(regexCache) < maxRegexCacheSize {
		regexCache[pattern] = re
	}
	regexCacheMu.Unlock()

	return re
}
