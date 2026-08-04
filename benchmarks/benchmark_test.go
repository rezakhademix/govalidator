package benchmarks

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	stdValidator "github.com/go-playground/validator/v10"
	govalidator "github.com/rezakhademix/govalidator/v2"
)

type UserCreateReq struct {
	FirstName     string `json:"first_name" validate:"required,min=1,max=100"`
	LastName      string `json:"last_name" validate:"required,min=1,max=100"`
	PhoneNumber   string `json:"phone_number" validate:"required,min=10,max=15"`
	Email         string `json:"email,omitempty" validate:"required,email"`
	FatherName    string `json:"father_name" validate:"required,min=1,max=100"`
	CertificateID string `json:"certificate_id" validate:"required,min=1,max=50"`
	BirthDate     string `json:"birth_date" validate:"required,datetime=2006-01-02"`
	CompanyID     int    `json:"company_id" validate:"required"`
	Gender        int8   `json:"gender" validate:"required,oneof=0 1"`
}

// std is created once: go-playground/validator caches struct metadata per
// instance, so creating it inside the benchmark loop would unfairly measure
// repeated cache rebuilding instead of validation.
var std = stdValidator.New()

// Standard library validation using go-playground/validator
func validateWithStdLibrary(req UserCreateReq) error {
	return std.Struct(req)
}

// Validation using ozzo-validation
func validateWithOzzo(req UserCreateReq) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.FirstName, validation.Required, validation.Length(1, 100)),
		validation.Field(&req.LastName, validation.Required, validation.Length(1, 100)),
		validation.Field(&req.PhoneNumber, validation.Required, validation.Length(10, 15)),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.FatherName, validation.Required, validation.Length(1, 100)),
		validation.Field(&req.CertificateID, validation.Required, validation.Length(1, 50)),
		validation.Field(&req.BirthDate, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&req.CompanyID, validation.Required),
		// values must be int8 to match the field type: ozzo compares
		// interface values, and int8(1) never equals int(1).
		validation.Field(&req.Gender, validation.Required, validation.In(int8(0), int8(1))),
	)
}

// Validation using rezakhademix/govalidator
func validateWithGoValidator(req UserCreateReq) map[string]string {
	v := govalidator.New()

	v.
		RequiredString(req.FirstName, "first_name", "First name is required").
		MinString(req.FirstName, 1, "first_name", "First name must be at least 1 character long").
		MaxString(req.FirstName, 100, "first_name", "First name cannot exceed 100 characters").
		RequiredString(req.LastName, "last_name", "Last name is required").
		MinString(req.LastName, 1, "last_name", "Last name must be at least 1 character long").
		MaxString(req.LastName, 100, "last_name", "Last name cannot exceed 100 characters").
		RequiredString(req.PhoneNumber, "phone_number", "Phone number is required").
		MinString(req.PhoneNumber, 10, "phone_number", "Phone number must be at least 10 characters long").
		MaxString(req.PhoneNumber, 15, "phone_number", "Phone number cannot exceed 15 characters").
		Email(req.Email, "email", "Invalid email format").
		RequiredString(req.FatherName, "father_name", "Father name is required").
		MinString(req.FatherName, 1, "father_name", "Father name must be at least 1 character long").
		MaxString(req.FatherName, 100, "father_name", "Father name cannot exceed 100 characters").
		RequiredString(req.CertificateID, "certificate_id", "Certificate ID is required").
		MinString(req.CertificateID, 1, "certificate_id", "Certificate ID must be at least 1 character long").
		MaxString(req.CertificateID, 50, "certificate_id", "Certificate ID cannot exceed 50 characters").
		RequiredString(req.BirthDate, "birth_date", "Birth date is required").
		Date("2006-01-02", req.BirthDate, "birth_date", "Invalid birth date format").
		RequiredInt(req.CompanyID, "company_id", "Company ID is required").
		RequiredInt(int(req.Gender), "gender", "Gender is required").
		MinInt(int(req.Gender), 0, "gender", "Gender must be at least 0").
		MaxInt(int(req.Gender), 1, "gender", "Gender must not exceed 1")

	if v.IsFailed() {
		return v.Errors()
	}

	return nil
}

func validReq() UserCreateReq {
	return UserCreateReq{
		FirstName:     "John",
		LastName:      "Doe",
		PhoneNumber:   "1234567890",
		Email:         "john.doe@example.com",
		FatherName:    "Michael",
		CertificateID: "12345",
		BirthDate:     "1990-01-01",
		CompanyID:     1,
		Gender:        1,
	}
}

// Benchmark tests
func BenchmarkStdLibrary(b *testing.B) {
	req := validReq()

	if err := validateWithStdLibrary(req); err != nil {
		b.Fatalf("expected passing input, got: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = validateWithStdLibrary(req)
	}
}

func BenchmarkOzzo(b *testing.B) {
	req := validReq()

	if err := validateWithOzzo(req); err != nil {
		b.Fatalf("expected passing input, got: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = validateWithOzzo(req)
	}
}

func BenchmarkGoValidator(b *testing.B) {
	req := validReq()

	if errs := validateWithGoValidator(req); len(errs) > 0 {
		b.Fatalf("expected passing input, got: %v", errs)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = validateWithGoValidator(req)
	}
}
