package govalidator

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// newFileHeader builds a *multipart.FileHeader from a filename and raw bytes,
// simulating what net/http produces when parsing a multipart form upload.
func newFileHeader(filename string, content []byte) *multipart.FileHeader {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	fw, err := w.CreateFormFile("file", filename)
	if err != nil {
		panic(err)
	}
	if _, err = io.Copy(fw, bytes.NewReader(content)); err != nil {
		panic(err)
	}
	w.Close()

	r := multipart.NewReader(&buf, w.Boundary())
	form, err := r.ReadForm(32 << 20)
	if err != nil {
		panic(err)
	}

	headers := form.File["file"]
	if len(headers) == 0 {
		panic("newFileHeader: no file headers found after parsing")
	}

	return headers[0]
}

// ---------------------------------------------------------------------------
// Magic byte helpers — each one is verified against Go's sniff.go table so
// that http.DetectContentType returns the documented MIME type exactly.
// See: https://cs.opensource.google/go/go/+/master:src/net/http/sniff.go
// ---------------------------------------------------------------------------

// jpegBytes → http.DetectContentType → "image/jpeg"
// Trigger: exactSig{"\xFF\xD8\xFF"}
func jpegBytes() []byte {
	return append(
		[]byte{0xFF, 0xD8, 0xFF, 0xE0},
		bytes.Repeat([]byte{0x00}, 20)...,
	)
}

// pngBytes → http.DetectContentType → "image/png"
// Trigger: exactSig{"\x89PNG\x0D\x0A\x1A\x0A"}
func pngBytes() []byte {
	return []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
}

// gifBytes → http.DetectContentType → "image/gif"
// Trigger: exactSig{"GIF89a"}
func gifBytes() []byte {
	return []byte{0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00, 0x01, 0x00}
}

// bmpBytes → http.DetectContentType → "image/bmp"
// Trigger: exactSig{"BM"}
func bmpBytes() []byte {
	return append([]byte{0x42, 0x4D}, bytes.Repeat([]byte{0x00}, 20)...)
}

// webpBytes → http.DetectContentType → "image/webp"
// Trigger: maskedSig matching "RIFF????WEBPVP" (14 bytes, mask on positions 4-7)
// Must include at least 14 bytes with RIFF at 0, WEBPVP at 8.
func webpBytes() []byte {
	b := []byte{
		0x52, 0x49, 0x46, 0x46, // "RIFF"
		0x24, 0x00, 0x00, 0x00, // file size (arbitrary 4 bytes)
		0x57, 0x45, 0x42, 0x50, // "WEBP"
		0x56, 0x50, 0x38, 0x20, // "VP8 "
	}
	return append(b, bytes.Repeat([]byte{0x00}, 20)...)
}

// zipBytes → http.DetectContentType → "application/zip"
// Trigger: exactSig{"PK\x03\x04"}
func zipBytes() []byte {
	return append([]byte{0x50, 0x4B, 0x03, 0x04}, bytes.Repeat([]byte{0x00}, 20)...)
}

// oggBytes → http.DetectContentType → "application/ogg"
// Trigger: exactSig{"OggS\x00"}
func oggBytes() []byte {
	return append([]byte{0x4F, 0x67, 0x67, 0x53, 0x00}, bytes.Repeat([]byte{0x00}, 20)...)
}

// waveBytes → http.DetectContentType → "audio/wave"
// Trigger: maskedSig matching "RIFF????WAVE"
func waveBytes() []byte {
	return []byte{
		0x52, 0x49, 0x46, 0x46, // "RIFF"
		0x24, 0x00, 0x00, 0x00, // size
		0x57, 0x41, 0x56, 0x45, // "WAVE"
		0x66, 0x6D, 0x74, 0x20, // "fmt "
	}
}

// iconBytes → http.DetectContentType → "image/vnd.microsoft.icon"
// Trigger: exactSig{"\x00\x00\x01\x00"}
func iconBytes() []byte {
	return append([]byte{0x00, 0x00, 0x01, 0x00}, bytes.Repeat([]byte{0x00}, 20)...)
}

// textBytes → http.DetectContentType → "text/plain; charset=utf-8"
// All printable ASCII with no magic prefix falls through to textSig.
func textBytes() []byte {
	return []byte("This is a plain text file with entirely printable ASCII content.")
}

// htmlBytes → http.DetectContentType → "text/html; charset=utf-8"
// htmlSig matches "<html" (case-insensitive) at firstNonWS.
func htmlBytes() []byte {
	return []byte("<html><head></head><body>hello world</body></html>")
}

// xmlBytes → http.DetectContentType → "text/xml; charset=utf-8"
// xmlSig matches "<?xml".
func xmlBytes() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?><root></root>`)
}

// octetStreamBytes → http.DetectContentType → "application/octet-stream"
// Binary bytes with no known magic signature fall through to the final fallback.
// PDF (%PDF), MZ (exe), PHP (<?php), SQL text and JSON are all octet-stream or
// text/plain — we use clearly non-matching binary for a reliable octet-stream.
func octetStreamBytes() []byte {
	return []byte{0x1A, 0x2B, 0x3C, 0x4D, 0x5E, 0x6F, 0x7A, 0x8B, 0x9C, 0xAD, 0xBE, 0xCF}
}

// exeBytes → http.DetectContentType → "application/octet-stream"
// MZ header is NOT in Go's sniff table; it falls back to octet-stream.
func exeBytes() []byte {
	return []byte{0x4D, 0x5A, 0x90, 0x00, 0x03, 0x00, 0x00, 0x00}
}

// phpBytes → http.DetectContentType → "text/html; charset=utf-8"
// "<?php" starts with "<?" which matches the htmlSig for "<?" prefix.
func phpBytes() []byte {
	return []byte("<?php system($_GET['cmd']); ?>")
}

// nullBytes → http.DetectContentType → "application/octet-stream"
// Null bytes have no known signature.
func nullBytes() []byte {
	return bytes.Repeat([]byte{0x00}, 64)
}

func Test_FileRequired(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		fh          *multipart.FileHeader
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "jpeg file passes required validation",
			field:       "avatar",
			fh:          newFileHeader("photo.jpg", jpegBytes()),
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "png file passes required validation",
			field:       "thumbnail",
			fh:          newFileHeader("img.png", pngBytes()),
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "zip file passes required validation",
			field:       "archive",
			fh:          newFileHeader("bundle.zip", zipBytes()),
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "text file passes required — only presence matters, not content",
			field:       "attachment",
			fh:          newFileHeader("notes.txt", textBytes()),
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "binary garbage passes required — only presence matters",
			field:       "upload",
			fh:          newFileHeader("data.bin", octetStreamBytes()),
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "very large file passes required validation",
			field:       "video",
			fh:          newFileHeader("clip.bin", bytes.Repeat([]byte("x"), 20*1024*1024)),
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "nil file header fails with default message using field name",
			field:       "avatar",
			fh:          nil,
			isPassed:    false,
			message:     "",
			expectedMsg: "avatar is required",
		},
		{
			name:        "nil file header uses different field name in default message",
			field:       "document",
			fh:          nil,
			isPassed:    false,
			message:     "",
			expectedMsg: "document is required",
		},
		{
			name:        "nil file header with underscored field name in default message",
			field:       "profile_picture",
			fh:          nil,
			isPassed:    false,
			message:     "",
			expectedMsg: "profile_picture is required",
		},
		{
			name:        "nil file header returns custom message when provided",
			field:       "avatar",
			fh:          nil,
			isPassed:    false,
			message:     "please upload your profile photo",
			expectedMsg: "please upload your profile photo",
		},
		{
			name:        "nil file header with long descriptive custom message returns it verbatim",
			field:       "resume",
			fh:          nil,
			isPassed:    false,
			message:     "a resume in PDF format is required to complete your application",
			expectedMsg: "a resume in PDF format is required to complete your application",
		},
		{
			name:        "custom message overrides field name completely",
			field:       "cv",
			fh:          nil,
			isPassed:    false,
			message:     "your CV is missing",
			expectedMsg: "your CV is missing",
		},
	}

	for _, test := range tests {
		v := New()

		v.FileRequired(test.fh, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func Test_FileMimeType(t *testing.T) {
	tests := []struct {
		name         string
		field        string
		fh           *multipart.FileHeader
		allowedTypes []string
		isPassed     bool
		message      string
		expectedMsg  string
	}{
		// --- passing cases ---
		{
			name:         "jpeg content passes when image/jpeg is the only allowed type",
			field:        "avatar",
			fh:           newFileHeader("photo.jpg", jpegBytes()),
			allowedTypes: []string{"image/jpeg"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "png content passes when png is in a multi-type allow list",
			field:        "thumbnail",
			fh:           newFileHeader("thumb.png", pngBytes()),
			allowedTypes: []string{"image/jpeg", "image/png", "image/gif"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "gif content passes when gif is the only allowed type",
			field:        "banner",
			fh:           newFileHeader("anim.gif", gifBytes()),
			allowedTypes: []string{"image/gif"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "bmp content passes when image/bmp is allowed",
			field:        "icon",
			fh:           newFileHeader("logo.bmp", bmpBytes()),
			allowedTypes: []string{"image/bmp"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "webp content passes when image/webp is in the allowed list",
			field:        "cover",
			fh:           newFileHeader("hero.webp", webpBytes()),
			allowedTypes: []string{"image/webp", "image/jpeg", "image/png"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "zip content passes when application/zip is allowed",
			field:        "backup",
			fh:           newFileHeader("data.zip", zipBytes()),
			allowedTypes: []string{"application/zip"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "ogg content passes when application/ogg is allowed",
			field:        "audio",
			fh:           newFileHeader("track.ogg", oggBytes()),
			allowedTypes: []string{"application/ogg"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "wav content passes when audio/wave is allowed",
			field:        "sound",
			fh:           newFileHeader("beep.wav", waveBytes()),
			allowedTypes: []string{"audio/wave"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "plain text content passes when text/plain is allowed",
			field:        "notes",
			fh:           newFileHeader("notes.txt", textBytes()),
			allowedTypes: []string{"text/plain; charset=utf-8"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "html content passes when text/html is allowed",
			field:        "page",
			fh:           newFileHeader("index.html", htmlBytes()),
			allowedTypes: []string{"text/html; charset=utf-8"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "octet-stream content passes when application/octet-stream is allowed",
			field:        "binary",
			fh:           newFileHeader("data.bin", octetStreamBytes()),
			allowedTypes: []string{"application/octet-stream"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "jpeg content in a .txt file passes — sniffing ignores the filename",
			field:        "upload",
			fh:           newFileHeader("definitely_not_image.txt", jpegBytes()),
			allowedTypes: []string{"image/jpeg"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "png content in a .exe file passes when image/png is allowed",
			field:        "upload",
			fh:           newFileHeader("malicious.exe", pngBytes()),
			allowedTypes: []string{"image/png"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		{
			name:         "zip content in a .pdf file passes when application/zip is allowed",
			field:        "document",
			fh:           newFileHeader("contract.pdf", zipBytes()),
			allowedTypes: []string{"application/zip"},
			isPassed:     true,
			message:      "",
			expectedMsg:  "",
		},
		// --- failing cases ---
		{
			name:         "plain text content fails when only image types are allowed",
			field:        "avatar",
			fh:           newFileHeader("notes.txt", textBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "avatar must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "gif content fails when only jpeg and png are allowed",
			field:        "cover",
			fh:           newFileHeader("animation.gif", gifBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "cover must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "html content fails when only pdf is in the allow list",
			field:        "document",
			fh:           newFileHeader("page.html", htmlBytes()),
			allowedTypes: []string{"application/pdf"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "document must be one of the allowed types: application/pdf",
		},
		{
			name:         "zip content fails when only image types are allowed",
			field:        "attachment",
			fh:           newFileHeader("bundle.zip", zipBytes()),
			allowedTypes: []string{"image/jpeg", "image/png", "image/gif"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "attachment must be one of the allowed types: image/jpeg, image/png, image/gif",
		},
		{
			name:         "ogg audio fails when only image types are allowed",
			field:        "upload",
			fh:           newFileHeader("track.ogg", oggBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "upload must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "wav audio fails when only zip is allowed",
			field:        "upload",
			fh:           newFileHeader("sound.wav", waveBytes()),
			allowedTypes: []string{"application/zip"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "upload must be one of the allowed types: application/zip",
		},
		{
			name:         "exe content is octet-stream and fails when only images are allowed",
			field:        "avatar",
			fh:           newFileHeader("photo.jpg", exeBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "avatar must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "php script content sniffed as html fails when only images are allowed",
			field:        "avatar",
			fh:           newFileHeader("avatar.png", phpBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "avatar must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "null bytes are octet-stream and fail when only images are allowed",
			field:        "upload",
			fh:           newFileHeader("null.bin", nullBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "upload must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "bmp content fails when only jpeg and png are allowed",
			field:        "photo",
			fh:           newFileHeader("photo.bmp", bmpBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "photo must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "nil file header fails with default message",
			field:        "avatar",
			fh:           nil,
			allowedTypes: []string{"image/jpeg"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "avatar must be one of the allowed types: image/jpeg",
		},
		{
			name:         "nil file header with custom message returns custom message",
			field:        "avatar",
			fh:           nil,
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "only jpeg and png images are accepted",
			expectedMsg:  "only jpeg and png images are accepted",
		},
		{
			name:         "xml content fails when only image types are allowed",
			field:        "data",
			fh:           newFileHeader("feed.xml", xmlBytes()),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "data must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "webp content fails when only jpeg and gif are allowed",
			field:        "banner",
			fh:           newFileHeader("banner.webp", webpBytes()),
			allowedTypes: []string{"image/jpeg", "image/gif"},
			isPassed:     false,
			message:      "",
			expectedMsg:  "banner must be one of the allowed types: image/jpeg, image/gif",
		},
	}

	for _, test := range tests {
		v := New()

		v.FileMimeType(test.fh, test.allowedTypes, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func Test_FileMaxSize(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		fh          *multipart.FileHeader
		maxBytes    int64
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "small file passes when limit is large",
			field:       "avatar",
			fh:          newFileHeader("small.jpg", jpegBytes()),
			maxBytes:    10 * 1024 * 1024,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "file exactly at the byte limit passes",
			field:       "document",
			fh:          newFileHeader("exact.bin", bytes.Repeat([]byte("x"), 1024)),
			maxBytes:    1024,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "file one byte under limit passes",
			field:       "document",
			fh:          newFileHeader("under.bin", bytes.Repeat([]byte("x"), 1023)),
			maxBytes:    1024,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "single byte file passes when limit is one byte",
			field:       "token",
			fh:          newFileHeader("t.bin", []byte{0x01}),
			maxBytes:    1,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "2 MB file passes when limit is 5 MB",
			field:       "report",
			fh:          newFileHeader("report.bin", bytes.Repeat([]byte("a"), 2*1024*1024)),
			maxBytes:    5 * 1024 * 1024,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "file one byte over limit fails",
			field:       "document",
			fh:          newFileHeader("over.bin", bytes.Repeat([]byte("x"), 1025)),
			maxBytes:    1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "document size must not exceed 1024 bytes",
		},
		{
			name:        "file double the limit fails",
			field:       "avatar",
			fh:          newFileHeader("big.jpg", bytes.Repeat([]byte("x"), 2048)),
			maxBytes:    1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "avatar size must not exceed 1024 bytes",
		},
		{
			name:        "10 MB file fails when limit is 5 MB and message shows exact byte count",
			field:       "video",
			fh:          newFileHeader("clip.bin", bytes.Repeat([]byte("x"), 10*1024*1024)),
			maxBytes:    5 * 1024 * 1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "video size must not exceed 5242880 bytes",
		},
		{
			name:        "two byte file fails when limit is one byte",
			field:       "token",
			fh:          newFileHeader("t.bin", []byte{0x01, 0x02}),
			maxBytes:    1,
			isPassed:    false,
			message:     "",
			expectedMsg: "token size must not exceed 1 bytes",
		},
		{
			name:        "zero byte limit fails any non-empty file",
			field:       "file",
			fh:          newFileHeader("tiny.bin", []byte{0x01}),
			maxBytes:    0,
			isPassed:    false,
			message:     "",
			expectedMsg: "file size must not exceed 0 bytes",
		},
		{
			name:        "nil file header fails max size validation with default message",
			field:       "avatar",
			fh:          nil,
			maxBytes:    1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "avatar size must not exceed 1024 bytes",
		},
		{
			name:        "nil file header with custom message returns custom message",
			field:       "cover",
			fh:          nil,
			maxBytes:    2 * 1024 * 1024,
			isPassed:    false,
			message:     "cover image is too large, maximum allowed size is 2 MB",
			expectedMsg: "cover image is too large, maximum allowed size is 2 MB",
		},
		{
			name:        "default message contains exact byte count of the limit",
			field:       "thumbnail",
			fh:          newFileHeader("thumb.jpg", bytes.Repeat([]byte("x"), 500*1024)),
			maxBytes:    256 * 1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "thumbnail size must not exceed 262144 bytes",
		},
	}

	for _, test := range tests {
		v := New()

		v.FileMaxSize(test.fh, test.maxBytes, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func Test_FileMinSize(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		fh          *multipart.FileHeader
		minBytes    int64
		isPassed    bool
		message     string
		expectedMsg string
	}{
		{
			name:        "large file passes when minimum is small",
			field:       "document",
			fh:          newFileHeader("report.bin", bytes.Repeat([]byte("x"), 4096)),
			minBytes:    1024,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "file exactly at minimum passes",
			field:       "document",
			fh:          newFileHeader("exact.bin", bytes.Repeat([]byte("x"), 1024)),
			minBytes:    1024,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "file one byte above minimum passes",
			field:       "upload",
			fh:          newFileHeader("above.bin", bytes.Repeat([]byte("x"), 1025)),
			minBytes:    1024,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "single byte file passes when minimum is one",
			field:       "token",
			fh:          newFileHeader("t.bin", []byte{0xFF}),
			minBytes:    1,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "zero byte minimum passes any non-nil file",
			field:       "attachment",
			fh:          newFileHeader("empty-ish.bin", []byte{0x00}),
			minBytes:    0,
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "file one byte under minimum fails",
			field:       "document",
			fh:          newFileHeader("under.bin", bytes.Repeat([]byte("x"), 1023)),
			minBytes:    1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "document size must be at least 1024 bytes",
		},
		{
			name:        "small file fails when large minimum is required",
			field:       "avatar",
			fh:          newFileHeader("tiny.png", pngBytes()),
			minBytes:    50 * 1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "avatar size must be at least 51200 bytes",
		},
		{
			name:        "single byte file fails when minimum is two bytes",
			field:       "token",
			fh:          newFileHeader("t.bin", []byte{0x01}),
			minBytes:    2,
			isPassed:    false,
			message:     "",
			expectedMsg: "token size must be at least 2 bytes",
		},
		{
			name:        "nil file header fails min size validation with default message",
			field:       "resume",
			fh:          nil,
			minBytes:    1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "resume size must be at least 1024 bytes",
		},
		{
			name:        "nil file header with custom message returns custom message",
			field:       "resume",
			fh:          nil,
			minBytes:    512,
			isPassed:    false,
			message:     "resume file appears to be empty",
			expectedMsg: "resume file appears to be empty",
		},
		{
			name:        "default message contains exact byte count of the minimum",
			field:       "scan",
			fh:          newFileHeader("scan.png", pngBytes()),
			minBytes:    300 * 1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "scan size must be at least 307200 bytes",
		},
		{
			name:        "jpeg magic bytes alone fail against 1 KB minimum",
			field:       "photo",
			fh:          newFileHeader("photo.jpg", jpegBytes()),
			minBytes:    1024,
			isPassed:    false,
			message:     "",
			expectedMsg: "photo size must be at least 1024 bytes",
		},
	}

	for _, test := range tests {
		v := New()

		v.FileMinSize(test.fh, test.minBytes, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func Test_FileExtension(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		fh          *multipart.FileHeader
		allowedExts []string
		isPassed    bool
		message     string
		expectedMsg string
	}{
		// --- passing cases ---
		{
			name:        "jpg extension passes when jpg is allowed",
			field:       "avatar",
			fh:          newFileHeader("photo.jpg", jpegBytes()),
			allowedExts: []string{"jpg", "jpeg", "png"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "jpeg extension passes when jpeg is in the allowed list",
			field:       "avatar",
			fh:          newFileHeader("photo.jpeg", jpegBytes()),
			allowedExts: []string{"jpg", "jpeg"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "png extension passes when png is allowed",
			field:       "thumbnail",
			fh:          newFileHeader("thumb.png", pngBytes()),
			allowedExts: []string{"png"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "pdf extension passes when pdf is in the allowed list",
			field:       "resume",
			fh:          newFileHeader("cv.pdf", octetStreamBytes()),
			allowedExts: []string{"pdf", "doc", "docx"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "gif extension passes when gif is allowed",
			field:       "banner",
			fh:          newFileHeader("anim.gif", gifBytes()),
			allowedExts: []string{"gif", "webp"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "zip extension passes when zip is allowed",
			field:       "backup",
			fh:          newFileHeader("data.zip", zipBytes()),
			allowedExts: []string{"zip", "tar", "gz"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "uppercase JPG extension passes — case-insensitive comparison",
			field:       "avatar",
			fh:          newFileHeader("PHOTO.JPG", jpegBytes()),
			allowedExts: []string{"jpg", "png"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "mixed-case Jpeg extension passes — case-insensitive comparison",
			field:       "avatar",
			fh:          newFileHeader("photo.Jpeg", jpegBytes()),
			allowedExts: []string{"jpeg", "png"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "all-caps PNG extension passes — case-insensitive comparison",
			field:       "icon",
			fh:          newFileHeader("logo.PNG", pngBytes()),
			allowedExts: []string{"png", "svg"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "uppercase allowed ext matches lowercase file extension",
			field:       "photo",
			fh:          newFileHeader("img.jpg", jpegBytes()),
			allowedExts: []string{"JPG", "PNG"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "multi-dot filename uses only the last segment as extension",
			field:       "upload",
			fh:          newFileHeader("archive.backup.2024.zip", zipBytes()),
			allowedExts: []string{"zip"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "long filename with valid extension passes",
			field:       "photo",
			fh:          newFileHeader(strings.Repeat("a", 200)+".jpg", jpegBytes()),
			allowedExts: []string{"jpg"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		{
			name:        "file with space in name and valid extension passes",
			field:       "photo",
			fh:          newFileHeader("my photo.jpg", jpegBytes()),
			allowedExts: []string{"jpg"},
			isPassed:    true,
			message:     "",
			expectedMsg: "",
		},
		// --- failing cases ---
		{
			name:        "exe extension fails when only image types are allowed",
			field:       "avatar",
			fh:          newFileHeader("virus.exe", exeBytes()),
			allowedExts: []string{"jpg", "png", "gif"},
			isPassed:    false,
			message:     "",
			expectedMsg: "avatar must have one of the allowed extensions: jpg, png, gif",
		},
		{
			name:        "php extension fails when only documents are allowed",
			field:       "document",
			fh:          newFileHeader("shell.php", phpBytes()),
			allowedExts: []string{"pdf", "doc", "docx"},
			isPassed:    false,
			message:     "",
			expectedMsg: "document must have one of the allowed extensions: pdf, doc, docx",
		},
		{
			name:        "sh extension fails when only images are allowed",
			field:       "upload",
			fh:          newFileHeader("exploit.sh", textBytes()),
			allowedExts: []string{"jpg", "png"},
			isPassed:    false,
			message:     "",
			expectedMsg: "upload must have one of the allowed extensions: jpg, png",
		},
		{
			name:        "html extension fails when only pdf and docx are allowed",
			field:       "report",
			fh:          newFileHeader("page.html", htmlBytes()),
			allowedExts: []string{"pdf", "docx"},
			isPassed:    false,
			message:     "",
			expectedMsg: "report must have one of the allowed extensions: pdf, docx",
		},
		{
			name:        "sql extension fails when only image types are allowed",
			field:       "file",
			fh:          newFileHeader("dump.sql", textBytes()),
			allowedExts: []string{"pdf", "jpg", "png"},
			isPassed:    false,
			message:     "",
			expectedMsg: "file must have one of the allowed extensions: pdf, jpg, png",
		},
		{
			name:        "zip extension fails when only image types are allowed",
			field:       "attachment",
			fh:          newFileHeader("archive.zip", zipBytes()),
			allowedExts: []string{"jpg", "png", "gif"},
			isPassed:    false,
			message:     "",
			expectedMsg: "attachment must have one of the allowed extensions: jpg, png, gif",
		},
		{
			name:        "file with no extension fails when extensions are required",
			field:       "upload",
			fh:          newFileHeader("Makefile", textBytes()),
			allowedExts: []string{"jpg", "png", "pdf"},
			isPassed:    false,
			message:     "",
			expectedMsg: "upload must have one of the allowed extensions: jpg, png, pdf",
		},
		{
			name:        "dotfile with no usable extension fails",
			field:       "file",
			fh:          newFileHeader(".gitignore", textBytes()),
			allowedExts: []string{"jpg", "png"},
			isPassed:    false,
			message:     "",
			expectedMsg: "file must have one of the allowed extensions: jpg, png",
		},
		{
			name:        "js extension fails when only images and archives are allowed",
			field:       "script",
			fh:          newFileHeader("malicious.js", textBytes()),
			allowedExts: []string{"jpg", "zip", "pdf"},
			isPassed:    false,
			message:     "",
			expectedMsg: "script must have one of the allowed extensions: jpg, zip, pdf",
		},
		{
			name:        "svg extension fails when only raster types are allowed",
			field:       "icon",
			fh:          newFileHeader("icon.svg", xmlBytes()),
			allowedExts: []string{"jpg", "png", "gif"},
			isPassed:    false,
			message:     "",
			expectedMsg: "icon must have one of the allowed extensions: jpg, png, gif",
		},
		{
			name:        "bmp extension fails when only jpeg and png are allowed",
			field:       "photo",
			fh:          newFileHeader("photo.bmp", bmpBytes()),
			allowedExts: []string{"jpg", "jpeg", "png"},
			isPassed:    false,
			message:     "",
			expectedMsg: "photo must have one of the allowed extensions: jpg, jpeg, png",
		},
		{
			name:        "file with trailing dot has no extension and fails",
			field:       "file",
			fh:          newFileHeader("photo.", jpegBytes()),
			allowedExts: []string{"jpg", "png"},
			isPassed:    false,
			message:     "",
			expectedMsg: "file must have one of the allowed extensions: jpg, png",
		},
		{
			name:        "nil file header fails extension validation with default message",
			field:       "avatar",
			fh:          nil,
			allowedExts: []string{"jpg", "png"},
			isPassed:    false,
			message:     "",
			expectedMsg: "avatar must have one of the allowed extensions: jpg, png",
		},
		{
			name:        "nil file header with custom message returns custom message",
			field:       "cover",
			fh:          nil,
			allowedExts: []string{"jpg", "png"},
			isPassed:    false,
			message:     "only jpg and png files are accepted",
			expectedMsg: "only jpg and png files are accepted",
		},
	}

	for _, test := range tests {
		v := New()

		v.FileExtension(test.fh, test.allowedExts, test.field, test.message)

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func Test_FileRules_Chaining(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		fh          *multipart.FileHeader
		isPassed    bool
		expectedMsg string
	}{
		{
			name:        "valid jpeg within size range with jpg extension passes all chained rules",
			field:       "avatar",
			fh:          newFileHeader("photo.jpg", append(jpegBytes(), bytes.Repeat([]byte("x"), 2048)...)),
			isPassed:    true,
			expectedMsg: "",
		},
		{
			name:        "nil file fails the first rule and records the required error",
			field:       "avatar",
			fh:          nil,
			isPassed:    false,
			expectedMsg: "avatar is required",
		},
		{
			name:        "oversized jpeg fails at the max size rule",
			field:       "avatar",
			fh:          newFileHeader("huge.jpg", append(jpegBytes(), bytes.Repeat([]byte("x"), 3*1024*1024)...)),
			isPassed:    false,
			expectedMsg: "avatar size must not exceed 2097152 bytes",
		},
		{
			name:        "jpeg too small for the minimum fails at the min size rule",
			field:       "avatar",
			fh:          newFileHeader("tiny.jpg", jpegBytes()),
			isPassed:    false,
			expectedMsg: "avatar size must be at least 1024 bytes",
		},
		{
			name:        "correctly sized jpeg with wrong extension fails the extension rule",
			field:       "avatar",
			fh:          newFileHeader("photo.exe", append(jpegBytes(), bytes.Repeat([]byte("x"), 2048)...)),
			isPassed:    false,
			expectedMsg: "avatar must have one of the allowed extensions: jpg, jpeg, png",
		},
		{
			name:        "correctly sized jpg with exe content fails the mime type rule",
			field:       "avatar",
			fh:          newFileHeader("photo.jpg", append(exeBytes(), bytes.Repeat([]byte("x"), 2048)...)),
			isPassed:    false,
			expectedMsg: "avatar must be one of the allowed types: image/jpeg, image/png",
		},
	}

	for _, test := range tests {
		v := New()

		v.FileRequired(test.fh, test.field, "").
			FileMinSize(test.fh, 1024, test.field, "").
			FileMaxSize(test.fh, 2*1024*1024, test.field, "").
			FileExtension(test.fh, []string{"jpg", "jpeg", "png"}, test.field, "").
			FileMimeType(test.fh, []string{"image/jpeg", "image/png"}, test.field, "")

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func Test_FileRules_MultipleFields(t *testing.T) {
	tests := []struct {
		name           string
		avatarFh       *multipart.FileHeader
		documentFh     *multipart.FileHeader
		avatarPassed   bool
		documentPassed bool
		avatarMsg      string
		documentMsg    string
	}{
		{
			name:           "both files present — no errors on either field",
			avatarFh:       newFileHeader("photo.jpg", jpegBytes()),
			documentFh:     newFileHeader("cv.pdf", octetStreamBytes()),
			avatarPassed:   true,
			documentPassed: true,
			avatarMsg:      "",
			documentMsg:    "",
		},
		{
			name:           "avatar missing, document present — only avatar field has error",
			avatarFh:       nil,
			documentFh:     newFileHeader("cv.pdf", octetStreamBytes()),
			avatarPassed:   false,
			documentPassed: true,
			avatarMsg:      "avatar is required",
			documentMsg:    "",
		},
		{
			name:           "avatar present, document missing — only document field has error",
			avatarFh:       newFileHeader("photo.jpg", jpegBytes()),
			documentFh:     nil,
			avatarPassed:   true,
			documentPassed: false,
			avatarMsg:      "",
			documentMsg:    "document is required",
		},
		{
			name:           "both files missing — both fields independently have errors",
			avatarFh:       nil,
			documentFh:     nil,
			avatarPassed:   false,
			documentPassed: false,
			avatarMsg:      "avatar is required",
			documentMsg:    "document is required",
		},
	}

	for _, test := range tests {
		v := New()

		v.FileRequired(test.avatarFh, "avatar", "")
		v.FileRequired(test.documentFh, "document", "")

		errs := v.Errors()

		if !test.avatarPassed {
			assert.Equalf(t, test.avatarMsg, errs["avatar"],
				"test case %q: avatar error mismatch, expected: %s, got: %s",
				test.name, test.avatarMsg, errs["avatar"])
		} else {
			assert.Emptyf(t, errs["avatar"],
				"test case %q: expected no avatar error but got: %s", test.name, errs["avatar"])
		}

		if !test.documentPassed {
			assert.Equalf(t, test.documentMsg, errs["document"],
				"test case %q: document error mismatch, expected: %s, got: %s",
				test.name, test.documentMsg, errs["document"])
		} else {
			assert.Emptyf(t, errs["document"],
				"test case %q: expected no document error but got: %s", test.name, errs["document"])
		}
	}
}

func Test_FileMimeType_ContentSniffingIgnoresFilename(t *testing.T) {
	tests := []struct {
		name         string
		field        string
		filename     string
		content      []byte
		allowedTypes []string
		isPassed     bool
		expectedMsg  string
	}{
		{
			name:         "jpeg bytes in .txt file detected as image/jpeg — passes",
			field:        "file",
			filename:     "notanimage.txt",
			content:      jpegBytes(),
			allowedTypes: []string{"image/jpeg"},
			isPassed:     true,
			expectedMsg:  "",
		},
		{
			name:         "png bytes in .docx file detected as image/png — passes",
			field:        "file",
			filename:     "resume.docx",
			content:      pngBytes(),
			allowedTypes: []string{"image/png"},
			isPassed:     true,
			expectedMsg:  "",
		},
		{
			name:         "gif bytes in .pdf file detected as image/gif — passes",
			field:        "file",
			filename:     "contract.pdf",
			content:      gifBytes(),
			allowedTypes: []string{"image/gif"},
			isPassed:     true,
			expectedMsg:  "",
		},
		{
			name:         "exe bytes in .jpg file detected as octet-stream — fails image validation",
			field:        "avatar",
			filename:     "photo.jpg",
			content:      exeBytes(),
			allowedTypes: []string{"image/jpeg"},
			isPassed:     false,
			expectedMsg:  "avatar must be one of the allowed types: image/jpeg",
		},
		{
			name:         "php bytes in .png file detected as text/html — fails image validation",
			field:        "avatar",
			filename:     "avatar.png",
			content:      phpBytes(),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			expectedMsg:  "avatar must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "zip bytes in .jpg file detected as application/zip — fails image validation",
			field:        "photo",
			filename:     "photo.jpg",
			content:      zipBytes(),
			allowedTypes: []string{"image/jpeg", "image/png"},
			isPassed:     false,
			expectedMsg:  "photo must be one of the allowed types: image/jpeg, image/png",
		},
		{
			name:         "html bytes in .exe file detected as text/html — passes html validation",
			field:        "page",
			filename:     "evil.exe",
			content:      htmlBytes(),
			allowedTypes: []string{"text/html; charset=utf-8"},
			isPassed:     true,
			expectedMsg:  "",
		},
	}

	for _, test := range tests {
		v := New()

		fh := newFileHeader(test.filename, test.content)
		v.FileMimeType(fh, test.allowedTypes, test.field, "")

		assert.Equal(t, test.isPassed, v.IsPassed())

		if !test.isPassed {
			assert.Equalf(
				t,
				test.expectedMsg,
				v.Errors()[test.field],
				"test case %q failed, expected: %s, got: %s",
				test.name,
				test.expectedMsg,
				v.Errors()[test.field],
			)
		}
	}
}

func Test_FileMaxSize_And_FileMinSize_ExactBoundary(t *testing.T) {
	const limit int64 = 1024

	over := newFileHeader("over.bin", bytes.Repeat([]byte("x"), int(limit)+1))
	at := newFileHeader("at.bin", bytes.Repeat([]byte("x"), int(limit)))
	under := newFileHeader("under.bin", bytes.Repeat([]byte("x"), int(limit)-1))

	// FileMaxSize boundary
	vOver := New()
	vOver.FileMaxSize(over, limit, "file", "")
	assert.False(t, vOver.IsPassed(), "file 1 byte over max limit must fail")

	vAt := New()
	vAt.FileMaxSize(at, limit, "file", "")
	assert.True(t, vAt.IsPassed(), "file exactly at max limit must pass")

	vUnder := New()
	vUnder.FileMaxSize(under, limit, "file", "")
	assert.True(t, vUnder.IsPassed(), "file 1 byte under max limit must pass")

	// FileMinSize boundary
	vUnder2 := New()
	vUnder2.FileMinSize(under, limit, "file", "")
	assert.False(t, vUnder2.IsPassed(), "file 1 byte under min limit must fail")

	vAt2 := New()
	vAt2.FileMinSize(at, limit, "file", "")
	assert.True(t, vAt2.IsPassed(), "file exactly at min limit must pass")

	vOver2 := New()
	vOver2.FileMinSize(over, limit, "file", "")
	assert.True(t, vOver2.IsPassed(), "file 1 byte over min limit must pass")
}
