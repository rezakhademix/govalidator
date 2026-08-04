package govalidator

import (
	"strings"
	"unicode/utf8"
)

const (
	// Len represents rule name which will be used to find the default error message.
	Len = "len"
	// LenList represents rule name which will be used to find the default error message.
	LenList = "lenList"
	// LenMsg is the default error message format for fields with Len validation rule.
	LenMsg = "%s should be %d characters"
	// LenListMsg is the default error message format for fields with LenList validation rule.
	LenListMsg = "%s should have %d items"
)

// LenString checks if the length of a string is equal to the given size or not.
//
// Example:
//
//	v := validator.New()
//	v.LenString("rez", 5, "username", "username must be 5 characters.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) LenString(s string, size int, field, msg string) Validator {
	if utf8.RuneCountInString(strings.TrimSpace(s)) != size {
		v.addError(field, v.msg(Len, msg, field, size))
	}

	return v
}

// LenInt checks if the length of the given integer is equal to the given size or not.
//
// Example:
//
//	v := validator.New()
//	v.LenInt(12345, 5, "zipcode", "Zip code must be 5 digits long.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) LenInt(i, size int, field, msg string) Validator {
	if intLen(i) != size {
		v.addError(field, v.msg(Len, msg, field, size))
	}

	return v
}

// LenSlice checks if the length of the given slice is equal to the given size or not.
//
// Example:
//
//	v := validator.New()
//	v.LenSlice([]int{1, 2, 3, 4, 5}, 5, "numbers", "the list must contain exactly 5 numbers.")
//	if v.IsFailed() {
//		 fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) LenSlice(s []any, size int, field, msg string) Validator {
	if len(s) != size {
		v.addError(field, v.msg(LenList, msg, field, size))
	}

	return v
}

// intLen mirrors len(strconv.Itoa(i)) — digit count plus one for a minus
// sign — without allocating a string.
func intLen(i int) int {
	n := 1
	if i < 0 {
		n++
	}

	for i /= 10; i != 0; i /= 10 {
		n++
	}

	return n
}
