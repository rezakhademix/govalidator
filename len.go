package govalidator

import (
	"strconv"
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
	v.check(utf8.RuneCountInString(strings.TrimSpace(s)) == size, field, v.msg(Len, msg, field, size))

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
	v.check(len(strconv.Itoa(i)) == size, field, v.msg(Len, msg, field, size))

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
	v.check(len(s) == size, field, v.msg(LenList, msg, field, size))

	return v
}
