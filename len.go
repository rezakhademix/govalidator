package validator

import (
	"strconv"
	"strings"
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
//	validator.LenString("rez", 5, "username", "username must be 5 characters.")
func (v *Validator) LenString(s string, size int, field, msg string) *Validator {
	v.Check(len(strings.TrimSpace(s)) == size, field, v.msg(Len, msg, field, size))

	return v
}

// LenInt checks if the length of the given integer is equal to the given size or not.
//
// Example:
//
//	validator.LenInt(12345, 5, "zipcode", "Zip code must be 5 digits long.")
func (v *Validator) LenInt(i, size int, field, msg string) *Validator {
	v.Check(len(strconv.Itoa(i)) == size, field, v.msg(Len, msg, field, size))

	return v
}

// LenSlice checks if the length of the given slice is equal to the given size or not.
//
// Example:
//
//	validator.LenSlice([]int{1, 2, 3, 4, 5}, 5, "numbers", "the list must contain exactly 5 numbers.")
func (v *Validator) LenSlice(s []any, size int, field, msg string) *Validator {
	v.Check(len(s) == size, field, v.msg(LenList, msg, field, size))

	return v
}
