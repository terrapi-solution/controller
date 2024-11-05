package errors

import (
	"errors"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewE(errors.New("error"), "message", INTERNAL)
	}
}

func BenchmarkNewInternal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewE(errors.New("error"), "message", INTERNAL)
	}
}

func BenchmarkError_Error(b *testing.B) {
	e := NewE(errors.New("error"), "message", INTERNAL)
	for i := 0; i < b.N; i++ {
		_ = e.Error()
	}
}

func BenchmarkError_HTTPStatusCode(b *testing.B) {
	e := NewE(errors.New("error"), "message", INTERNAL)
	for i := 0; i < b.N; i++ {
		_ = e.HTTPStatusCode()
	}
}
