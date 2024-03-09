package validator

import (
	"strconv"
	"strings"
)

const (
	// Len represents the rule name which will be used to find the default error message.
	Len = "len"
	// LenList represents the rule name which will be used to find the default error message.
	LenList = "lenList"
	// LenMsg is the default error message format for Len validation rule.
	LenMsg = "%s should be %d characters"
	// LenListMsg is the default error message format for LenSlice validation rule.
	LenListMsg = "%s should have %d items"
)

// LenString checks if length of a string equal to give size or not.
func (v *Validator) LenString(s string, size int, field, msg string) *Validator {
	v.Check(len(strings.TrimSpace(s)) == size, field, v.msg(Len, msg, field, size))

	return v
}

// LenInt checks if length of an integer is equal to given size or not.
func (v *Validator) LenInt(i, size int, field, msg string) *Validator {
	v.Check(len(strconv.Itoa(i)) == size, field, v.msg(Len, msg, field, size))

	return v
}

// LenSlice checks if length of a slice is equal to given size or not.
func (v *Validator) LenSlice(s []any, size int, field, msg string) *Validator {
	v.Check(len(s) == size, field, v.msg(LenList, msg, field, size))

	return v
}
