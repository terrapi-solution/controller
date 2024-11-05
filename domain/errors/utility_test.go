package errors

import (
	"reflect"
	"testing"
)

func TestNewInternal(t *testing.T) {
	got := NewInternal(nil, "message", "op")
	want := INTERNAL
	if !reflect.DeepEqual(want, got.Code) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestNewConflict(t *testing.T) {
	got := NewConflict(nil, "message", "op")
	want := CONFLICT
	if !reflect.DeepEqual(want, got.Code) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestNewInvalid(t *testing.T) {
	got := NewInvalid(nil, "message", "op")
	want := INVALID
	if !reflect.DeepEqual(want, got.Code) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestNewNotFound(t *testing.T) {
	got := NewNotFound(nil, "message", "op")
	want := NOTFOUND
	if !reflect.DeepEqual(want, got.Code) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestNewUnknown(t *testing.T) {
	got := NewUnknown(nil, "message", "op")
	want := UNKNOWN
	if !reflect.DeepEqual(want, got.Code) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestNewUnauthorized(t *testing.T) {
	got := NewUnauthorized(nil, "message", "op")
	want := UNAUTHORIZED
	if !reflect.DeepEqual(want, got.Code) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestNewForbidden(t *testing.T) {
	got := NewForbidden(nil, "message", "op")
	want := FORBIDDEN
	if !reflect.DeepEqual(want, got.Code) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}
