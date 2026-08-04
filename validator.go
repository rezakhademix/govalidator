// Package govalidator provides configurable rules for validating data of various types.
package govalidator

import (
	"errors"
	"fmt"
)

type (
	// Repository represent a repository for using in rules that needs a database connection to
	// check a record exists on database or not.
	Repository interface {
		// Exists checks whether a given value exists in a specified column of a table.
		// This function is typically used to validate the uniqueness of a value in a database.
		//
		// Parameters:
		// - value: The value to check for existence.
		// - table: The name of the table to search within.
		// - column: The column to check for the value.
		//
		// Returns:
		// - bool: Returns true if the value exists in the specified column of the table;
		//         otherwise, returns false.
		Exists(value any, table, column string) bool

		// ExistsExceptSelf checks whether a given value exists in a specified column of a table,
		// excluding a row identified by selfID. This is typically used to ensure the uniqueness
		// of a value in a table, while allowing updates to a row without triggering a false positive.
		//
		// Parameters:
		// - value: The value to check for existence.
		// - table: The name of the table to search within.
		// - column: The column to check for the value.
		// - selfID: The ID of the row to exclude from the check.
		//
		// Returns:
		// - bool: Returns true if the value exists in the specified column of the table, excluding
		//
		ExistsExceptSelf(value any, table, column string, selfID int) bool
	}

	// errBag holds the validation errors behind a pointer so that Validator
	// methods can keep value receivers while sharing one error store across
	// the whole method chain. The map itself is created lazily on the first
	// failure, keeping a fully passing validation free of map allocations.
	errBag struct {
		errs map[string]string
	}

	// Validator represents the validator structure
	Validator struct {
		bag  *errBag
		repo Repository
	}
)

var (
	// methodToErrorMessage contains each validation method and its corresponding error message.
	methodToErrorMessage = map[string]string{
		Required:      RequiredMsg,
		Exists:        ExistsMsg,
		Len:           LenMsg,
		LenList:       LenListMsg,
		Max:           MaxMsg,
		MaxString:     MaxStringMsg,
		MinString:     MinStringMsg,
		Min:           MinMsg,
		Between:       BetweenMsg,
		NotExists:     NotExistsMsg,
		Regex:         RegexMsg,
		Email:         EmailMsg,
		UUID:          UUIDMsg,
		Date:          DateMsg,
		Time:          TimeMsg,
		URL:           URLMsg,
		Before:        BeforeMsg,
		After:         AfterMsg,
		IP4:           IP4Msg,
		JSON:          JSONMsg,
		NumericString: NumericStringMsg,
		MAC:           MACMsg,
		RequiredFile:  RequiredFileMsg,
		FileMimeType:  FileMimeTypeMsg,
		FileMaxSize:   FileMaxSizeMsg,
		FileMinSize:   FileMinSizeMsg,
		FileExtension: FileExtensionMsg,
	}

	// ErrMethodMessageNotFound is the default message when a method does not have any error message on methodToErrorMessage.
	ErrMethodMessageNotFound = errors.New("method default validation message does not exist in methodToErrorMessage")
)

// New will return a new validator
func New() Validator {
	return Validator{
		bag: &errBag{},
	}
}

// WithRepo sets the desired repository for use in the Exists validation rule.
//
// Example:
//
//	validator := New().WithRepo(myRepository)
func (v Validator) WithRepo(r Repository) Validator {
	v.repo = r

	return v
}

// IsPassed checks if the validator result has passed or not.
func (v Validator) IsPassed() bool {
	return v.bag == nil || len(v.bag.errs) == 0
}

// IsFailed  checks if the validator result has failed or not.
func (v Validator) IsFailed() bool {
	return !v.IsPassed()
}

// Errors returns a map of all validator rule errors. When no rule has failed
// it returns a fresh empty map instead of materializing the shared store, so
// Errors stays a pure read and is safe to call from concurrent readers.
func (v Validator) Errors() map[string]string {
	if v.bag == nil || v.bag.errs == nil {
		return map[string]string{}
	}

	return v.bag.errs
}

// addError fills the errors map and prevents duplicate fields from being added to validator errors.
func (v Validator) addError(field, msg string) {
	if v.bag.errs == nil {
		v.bag.errs = make(map[string]string)
	}

	if _, exists := v.bag.errs[field]; !exists {
		v.bag.errs[field] = msg
	}
}

// msg returns the error message. If a custom error message is set, it returns the formatted custom message;
// otherwise, it returns the default message for the rule which has been set on the validator.
func (v Validator) msg(method, msg string, fieldArgs ...any) string {
	if msg != "" {
		return msg
	}

	defaultMsg, ok := methodToErrorMessage[method]
	if !ok {
		panic(ErrMethodMessageNotFound)
	}

	return fmt.Sprintf(defaultMsg, fieldArgs...)
}
