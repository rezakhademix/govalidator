//go:build !race

package govalidator

import "testing"

// Test_PassingValidationAllocsAtMostOnce mechanically guards the lazy
// error-message invariant: a fully passing validation must never format
// messages, so it performs at most one allocation (the shared error store).
// Excluded under -race because the race detector inflates allocation counts.
func Test_PassingValidationAllocsAtMostOnce(t *testing.T) {
	if testing.CoverMode() != "" {
		t.Skip("coverage instrumentation skews allocation counts")
	}

	req := UserCreateReq{
		FirstName:     "Reza",
		LastName:      "Khademi",
		PhoneNumber:   "09121234567",
		Email:         "rezakhademix@gmail.com",
		FatherName:    "Ali",
		CertificateID: "1234567890",
		BirthDate:     "1990-01-01",
		CompanyID:     42,
		Gender:        1,
	}

	allocs := testing.AllocsPerRun(200, func() {
		v := validateUserCreateReq(req)
		if v.IsFailed() {
			t.Fatalf("expected validation to pass, got errors: %v", v.Errors())
		}
	})

	if allocs > 1 {
		t.Errorf("passing validation made %v allocations per run, want at most 1 — a rule is likely formatting its error message eagerly", allocs)
	}
}
