// Package validator provides configurable rules for validating data of various types.
package validator

import (
	"errors"
	"fmt"
	"maps"
)

type (
	// Err is the defined type which will be returned when one or many validator rules fail.
	Err = map[string]string
	// Validator represents the validator structure
	Validator struct{}
)

var (
	// errs is a map of errors which has all of failed rule messages.
	errs = make(Err)

	// methodToErrorMessage contains each validation method and its corresponding error message.
	methodToErrorMessage = map[string]string{
		Required: RequiredMsg,
		Len:      LenMsg,
		Min:      MinMsg,
		Between:  BetweenMsg,
	}

	// ErrMethodMessageNotFound is the default message when a method does not have any error message on methodToErrorMessage.
	ErrMethodMessageNotFound = errors.New("method message does not exist")
)

// New will return a new validator
func New() *Validator {
	return &Validator{}
}

// IsPassed checks validator result is passed or not.
func (v *Validator) IsPassed() bool {
	return len(errs) == 0
}

// IsFailed checks validator result is failed or not.
func (v *Validator) IsFailed() bool {
	return !v.IsPassed()
}

// Errors returns a map of all validator rule errors.
func (v *Validator) Errors() Err {
	vErrs := maps.Clone(errs)

	errs = make(map[string]string)

	return vErrs
}

// Check is a dynamic method to define any custom validator rule by passing rule as a func or expression
// which will return a boolean.
func (v *Validator) Check(ok bool, field, msg string) {
	if !ok {
		v.addError(field, msg)
	}
}

// addErrors fills errors map and prevent duplicates field from being added to validator errors
func (v *Validator) addError(field, msg string) {
	if _, exists := errs[field]; !exists {
		errs[field] = msg
	}
}

// msg return error message and check if custom error message is set return formatted custom message
// otherwise return rule default message
func (v *Validator) msg(method, msg string, fieldArgs ...any) string {
	if msg != "" {
		return msg
	}

	defaultMsg, ok := methodToErrorMessage[method]
	if !ok {
		panic(ErrMethodMessageNotFound)
	}

	return fmt.Sprintf(defaultMsg, fieldArgs...)
}
