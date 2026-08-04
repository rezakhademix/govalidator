package govalidator

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_msg(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		msg         string
		expectedMsg string
	}{
		{
			name:        "test not exists method will result a panic",
			method:      "qwert",
			msg:         "",
			expectedMsg: "method default validation message does not exist in methodToErrorMessage",
		},
		{
			name:        "test empty string method will result a panic",
			method:      "",
			msg:         "",
			expectedMsg: "method default validation message does not exist in methodToErrorMessage",
		},
		{
			name:        "test empty space string method will result a panic",
			method:      " ",
			msg:         "",
			expectedMsg: "method default validation message does not exist in methodToErrorMessage",
		},
	}

	for _, test := range tests {
		v := New()

		assert.PanicsWithError(t, test.expectedMsg, func() { v.msg(test.method, test.msg) })
	}
}

func Test_ZeroValueValidatorReadsAreSafe(t *testing.T) {
	var v Validator

	assert.NotPanics(t, func() {
		assert.True(t, v.IsPassed())
		assert.False(t, v.IsFailed())
		assert.Empty(t, v.Errors())
	})
}

func Test_ErrorsIsSafeForConcurrentReads(t *testing.T) {
	v := New()
	v.RequiredString("reza", "name", "")

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 200; j++ {
				_ = v.Errors()
				_ = v.IsPassed()
			}
		}()
	}
	wg.Wait()
}
