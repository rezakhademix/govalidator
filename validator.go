package validator

import "maps"

type (
	Err       = map[string]string
	Validator struct{}
)

var errs = make(Err)

func New() *Validator {
	return &Validator{}
}

func (v *Validator) IsPassed() bool {
	return len(errs) == 0
}

func (v *Validator) IsFailed() bool {
	return !v.IsPassed()
}

func (v *Validator) Errors() Err {
	vErrs := maps.Clone(errs)

	errs = make(map[string]string)

	return vErrs
}

func (v *Validator) Check(ok bool, field, msg string) {
	if !ok {
		v.addErrors(field, msg)
	}
}

func (v *Validator) addErrors(field, msg string) {
	if _, exists := errs[field]; !exists {
		errs[field] = msg
	}
}
