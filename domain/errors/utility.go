package errors

import (
	"runtime"
	"strconv"
)

// newError is an alias for New by creating the pcs
// file line and constructing the error message.
func newError(err error, message, code, op string) *Error {
	_, file, line, _ := runtime.Caller(2)
	pcs := make([]uintptr, 100)
	_ = runtime.Callers(3, pcs)
	return &Error{
		Code:      code,
		Message:   message,
		Operation: op,
		Err:       err,
		fileLine:  file + ":" + strconv.Itoa(line),
		pcs:       pcs,
	}
}

// NewInternal returns an Error with a INTERNAL error code.
func NewInternal(err error, message, op string) *Error {
	return newError(err, message, INTERNAL, op)
}

// NewConflict returns an Error with a CONFLICT error code.
func NewConflict(err error, message, op string) *Error {
	return newError(err, message, CONFLICT, op)
}

// NewInvalid returns an Error with a INVALID error code.
func NewInvalid(err error, message, op string) *Error {
	return newError(err, message, INVALID, op)
}

// NewNotFound returns an Error with a NOTFOUND error code.
func NewNotFound(err error, message, op string) *Error {
	return newError(err, message, NOTFOUND, op)
}

// NewUnknown returns an Error with a UNKNOWN error code.
func NewUnknown(err error, message, op string) *Error {
	return newError(err, message, UNKNOWN, op)
}

// NewUnauthorized returns an Error with a UNAUTHORIZED error code.
func NewUnauthorized(err error, message, op string) *Error {
	return newError(err, message, UNAUTHORIZED, op)
}

// NewForbidden returns an Error with a FORBIDDEN error code.
func NewForbidden(err error, message, op string) *Error {
	return newError(err, message, FORBIDDEN, op)
}
