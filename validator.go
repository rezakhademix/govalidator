// Package validator provides configurable rules for validating data of various types.
package validator

import (
	"errors"
	"fmt"
	"maps"
)

type (
	// Err the defined type which will be returned when one or many validator rules failed.
	Err = map[string]string
	// Validator represents the validator structure
	Validator struct{}
)

var (
	// errs initiates map errors which has map[string]string type
	errs = make(Err)

	methodToErrorMessage = map[string]string{
		RequiredMethod: RequiredErrorMessage,
	}

	// ErrMethodMessageNotFound return error when method name doesn`t have default error message
	ErrMethodMessageNotFound = errors.New("rule error message not exists")
)

// New will return a new validator struct
func New() *Validator {
	return &Validator{}
}

// IsPassed is the method to check validation result is passed or not
func (v *Validator) IsPassed() bool {
	return len(errs) == 0
}

// IsFailed is the method to check validation result is failed or not
func (v *Validator) IsFailed() bool {
	return !v.IsPassed()
}

// Errors returns a map of errors of type map[string]string to
func (v *Validator) Errors() Err {
	vErrs := maps.Clone(errs)

	errs = make(map[string]string)

	return vErrs
}

// Check whether to add err to errors map or not.
// This check will be based on ok bool which will be considered as the rule to check
func (v *Validator) Check(ok bool, field, msg string) {
	if !ok {
		v.addErrors(field, msg)
	}
}

// addErrors fill errors map and prevent duplicates field from being added to validator errors
func (v *Validator) addErrors(field, msg string) {
	if _, exists := errs[field]; !exists {
		errs[field] = msg
	}
}

// FindErrorMessage return error message and check if custom error message is set return formatted custom message
// otherwise return rule default message
func FindErrorMessage(method, field string, msgArgs ...any) string {
	if len(msgArgs) == 1 {
		return msgArgs[0].(string)
	}

	if len(msgArgs) > 1 {
		return fmt.Sprintf(msgArgs[0].(string), msgArgs[1:])
	}

	format, ok := methodToErrorMessage[method]
	if !ok {
		panic(ErrMethodMessageNotFound)
	}

	return fmt.Sprintf(format, field)
}
