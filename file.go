package govalidator

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	// FileRequired represents rule name which will be used to find the default error message.
	FileRequired = "fileRequired"
	// FileRequiredMsg is the default error message format for fields with FileRequired validation rule.
	FileRequiredMsg = "%s is required"

	// FileMimeType represents rule name which will be used to find the default error message.
	FileMimeType = "fileMimeType"
	// FileMimeTypeMsg is the default error message format for fields with FileMimeType validation rule.
	FileMimeTypeMsg = "%s must be one of the allowed types: %s"

	// FileMaxSize represents rule name which will be used to find the default error message.
	FileMaxSize = "fileMaxSize"
	// FileMaxSizeMsg is the default error message format for fields with FileMaxSize validation rule.
	FileMaxSizeMsg = "%s size must not exceed %d bytes"

	// FileMinSize represents rule name which will be used to find the default error message.
	FileMinSize = "fileMinSize"
	// FileMinSizeMsg is the default error message format for fields with FileMinSize validation rule.
	FileMinSizeMsg = "%s size must be at least %d bytes"

	// FileExtension represents rule name which will be used to find the default error message.
	FileExtension = "fileExtension"
	// FileExtensionMsg is the default error message format for fields with FileExtension validation rule.
	FileExtensionMsg = "%s must have one of the allowed extensions: %s"
)

// FileRequired checks if a multipart file header is present and non-nil.
// A nil fh or a zero-size file means no file was uploaded for this field.
//
// Example:
//
//	_, fh, _ := r.FormFile("avatar")
//	v := validator.New()
//	v.FileRequired(fh, "avatar", "avatar is required.")
//	if v.IsFailed() {
//	    fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) FileRequired(fh *multipart.FileHeader, field, msg string) Validator {
	v.check(fh != nil && fh.Size > 0, field, v.msg(FileRequired, msg, field))

	return v
}

// FileMimeType checks whether the uploaded file's MIME type is among the allowed types.
// It sniffs the actual file content using net/http.DetectContentType (reads first 512 bytes),
// rather than trusting the client-supplied Content-Type header which can be spoofed.
//
// allowedTypes should be a list of MIME type strings, e.g. []string{"image/jpeg", "image/png"}.
//
// Example:
//
//	v := validator.New()
//	v.FileMimeType(fh, []string{"image/jpeg", "image/png"}, "avatar", "only JPEG and PNG images are allowed.")
//	if v.IsFailed() {
//	    fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) FileMimeType(fh *multipart.FileHeader, allowedTypes []string, field, msg string) Validator {
	v.check(isMimeTypeAllowed(fh, allowedTypes), field, v.msg(FileMimeType, msg, field, strings.Join(allowedTypes, ", ")))

	return v
}

// FileMaxSize checks that the uploaded file does not exceed maxBytes in size.
//
// Example:
//
//	v := validator.New()
//	v.FileMaxSize(fh, 5*1024*1024, "document", "document must not exceed 5 MB.")
//	if v.IsFailed() {
//	    fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) FileMaxSize(fh *multipart.FileHeader, maxBytes int64, field, msg string) Validator {
	v.check(fh != nil && fh.Size <= maxBytes, field, v.msg(FileMaxSize, msg, field, maxBytes))

	return v
}

// FileMinSize checks that the uploaded file meets the minimum size requirement in bytes.
//
// Example:
//
//	v := validator.New()
//	v.FileMinSize(fh, 1024, "document", "document must be at least 1 KB.")
//	if v.IsFailed() {
//	    fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) FileMinSize(fh *multipart.FileHeader, minBytes int64, field, msg string) Validator {
	v.check(fh != nil && fh.Size >= minBytes, field, v.msg(FileMinSize, msg, field, minBytes))

	return v
}

// FileExtension checks that the uploaded file's extension (derived from its filename)
// is among the allowed extensions. Extensions are compared case-insensitively and
// should be provided without the leading dot, e.g. []string{"jpg", "png", "pdf"}.
//
// Note: extension checking alone is not a security measure. Prefer FileMimeType for
// content-level validation, or combine both for defense in depth.
//
// Example:
//
//	v := validator.New()
//	v.FileExtension(fh, []string{"jpg", "jpeg", "png"}, "avatar", "only JPG and PNG files are allowed.")
//	if v.IsFailed() {
//	    fmt.Printf("validation errors: %#v\n", v.Errors())
//	}
func (v Validator) FileExtension(fh *multipart.FileHeader, allowedExts []string, field, msg string) Validator {
	v.check(isExtensionAllowed(fh, allowedExts), field, v.msg(FileExtension, msg, field, strings.Join(allowedExts, ", ")))

	return v
}

// isMimeTypeAllowed opens the file header, reads up to 512 bytes, and uses
// net/http.DetectContentType to determine the actual MIME type, then checks
// it against the list of allowed types.
func isMimeTypeAllowed(fh *multipart.FileHeader, allowedTypes []string) bool {
	if fh == nil {
		return false
	}

	f, err := fh.Open()
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil && n == 0 {
		return false
	}

	detected := http.DetectContentType(buf[:n])

	for _, allowed := range allowedTypes {
		if strings.EqualFold(detected, allowed) {
			return true
		}
	}

	return false
}

// isExtensionAllowed derives the extension from the file header's filename and
// checks it case-insensitively against the list of allowed extensions.
func isExtensionAllowed(fh *multipart.FileHeader, allowedExts []string) bool {
	if fh == nil {
		return false
	}

	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(fh.Filename)), ".")

	for _, allowed := range allowedExts {
		if strings.ToLower(allowed) == ext {
			return true
		}
	}

	return false
}
