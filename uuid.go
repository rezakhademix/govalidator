package govalidator

import "github.com/google/uuid"

const (
	// UUID represents rule name which will be used to find the default error message.
	UUID = "uuid"
	// UUIDMsg is the default error message format for fields with UUID validation rule.
	UUIDMsg = "%s is not a valid UUID"
)

// UUID validates that the field under validation is a valid RFC 4122 universally unique identifier (UUID).
// It accepts non-standard strings such as raw hex encoding xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// and 38 byte "Microsoft style" encodings, e.g., {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}.
//
// Example:
//
//	validator.UUID("f47ac10b-58cc-4372-a567-0e02b2c3d479", "uuid", "Invalid UUID format.")
func (v *Validator) UUID(u, field, msg string) *Validator {
	_, err := uuid.Parse(u)
	if err != nil {
		v.Check(false, field, v.msg(UUID, msg, field))

		return v
	}

	return v
}
