// Package govalidator provides configurable rules for validating data of various types.
package govalidator

import (
	"errors"
	"fmt"
	"maps"
)

type (
	// Err is the defined type which will be returned when one or many validator rules fail.
	Err = map[string]string
	// Validator represents the validator structure
	Validator struct {
		repo Repository
	}

	// Repository represent a repository for using in rules that needs a database connection to
	// check a record exists on database or not.
	Repository interface {
		Exists(value any, table, column string) bool
	}
)

var (
	// errs is a map of errors which has all of failed rule messages.
	errs = make(Err)

	// methodToErrorMessage contains each validation method and its corresponding error message.
	methodToErrorMessage = map[string]string{
		Required:  RequiredMsg,
		Exists:    ExistsMsg,
		Len:       LenMsg,
		LenList:   LenListMsg,
		Max:       MaxMsg,
		MaxString: MaxStringMsg,
		MinString: MinStringMsg,
		Min:       MinMsg,
		Between:   BetweenMsg,
		NotExists: NotExistsMsg,
		Regex:     RegexMsg,
		Email:     EmailMsg,
		UUID:      UUIDMsg,
		Date:      DateMsg,
	}

	// ErrMethodMessageNotFound is the default message when a method does not have any error message on methodToErrorMessage.
	ErrMethodMessageNotFound = errors.New("method default validation message does not exist in methodToErrorMessage")
)

// New will return a new validator
func New() *Validator {
	return &Validator{}
}

// WithRepo sets the desired repository for use in the Exists validation rule.
//
// Example:
//
//	validator := New().WithRepo(myRepository)
func (v *Validator) WithRepo(r Repository) *Validator {
	v.repo = r

	return v
}

// IsPassed checks if the validator result has passed or not.
func (v *Validator) IsPassed() bool {
	return len(errs) == 0
}

// IsFailed  checks if the validator result has failed or not.
func (v *Validator) IsFailed() bool {
	return !v.IsPassed()
}

// Errors returns a map of all validator rule errors.
func (v *Validator) Errors() Err {
	vErrs := maps.Clone(errs)

	errs = make(map[string]string)

	return vErrs
}

// Check is a dynamic method to define any custom validator rule by passing a rule as a function or expression
// which will return a boolean.
func (v *Validator) Check(ok bool, field, msg string) {
	if !ok {
		v.addError(field, msg)
	}
}

// addError fills the errors map and prevents duplicate fields from being added to validator errors.
func (v *Validator) addError(field, msg string) {
	if _, exists := errs[field]; !exists {
		errs[field] = msg
	}
}

// msg returns the error message. If a custom error message is set, it returns the formatted custom message;
// otherwise, it returns the default message for the rule which has been set on the validator.
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
