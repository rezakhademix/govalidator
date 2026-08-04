package govalidator

import (
	"testing"
)

// UserCreateReq is the sample DTO used in the README benchmarks.
type UserCreateReq struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email,omitempty"`
	FatherName    string `json:"father_name"`
	CertificateID string `json:"certificate_id"`
	BirthDate     string `json:"birth_date"`
	CompanyID     int    `json:"company_id"`
	Gender        int8   `json:"gender"`
}

func validateUserCreateReq(req UserCreateReq) Validator {
	v := New()

	v.RequiredString(req.FirstName, "first_name", "").
		MinString(req.FirstName, 2, "first_name", "").
		MaxString(req.FirstName, 50, "first_name", "").
		RequiredString(req.LastName, "last_name", "").
		MaxString(req.LastName, 50, "last_name", "").
		NumericString(req.PhoneNumber, "phone_number", "").
		Email(req.Email, "email", "").
		RequiredString(req.FatherName, "father_name", "").
		LenString(req.CertificateID, 10, "certificate_id", "").
		Date("2006-01-02", req.BirthDate, "birth_date", "").
		RequiredInt(req.CompanyID, "company_id", "").
		BetweenInt(int(req.Gender), 0, 2, "gender", "")

	return v
}

func Benchmark_UserCreateReq_Passing(b *testing.B) {
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

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v := validateUserCreateReq(req)
		if v.IsFailed() {
			b.Fatalf("expected validation to pass, got errors: %v", v.Errors())
		}
	}
}

func Benchmark_UserCreateReq_Failing(b *testing.B) {
	req := UserCreateReq{
		FirstName:     "R",
		LastName:      "",
		PhoneNumber:   "not-a-number",
		Email:         "invalid-email",
		FatherName:    "",
		CertificateID: "123",
		BirthDate:     "01-01-1990",
		CompanyID:     0,
		Gender:        9,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		v := validateUserCreateReq(req)
		if v.IsPassed() {
			b.Fatal("expected validation to fail")
		}
	}
}
