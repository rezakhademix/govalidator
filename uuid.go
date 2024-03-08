package validator

import "github.com/google/uuid"

const (
	// UUID represents the rule name which will be used to find the default error message.
	UUID = "uuid"
	// UUIDMsg is the default error message format for fields with the UUID validation rule.
	UUIDMsg = "%s is not a valid UUID"
)

// UUID The field under validation must be a valid RFC 4122 universally unique identifier (UUID).
// In addition, UUID accepts non-standard strings such as the raw hex encoding xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// and 38 byte "Microsoft style" encodings, e.g. {xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}.
func (v *Validator) UUID(u, field, msg string) *Validator {
	_, err := uuid.Parse(u)
	if err != nil {
		v.Check(false, field, v.msg(UUID, msg, field))

		return v
	}

	return v
}
