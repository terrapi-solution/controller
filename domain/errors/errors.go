package errors

import (
	"bytes"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

// Application error codes.
const (
	// CONFLICT - An action cannot be performed.
	CONFLICT = "conflict"
	// INTERNAL - Error within the application.
	INTERNAL = "internal"
	// INVALID - Validation failed.
	INVALID = "invalid"
	// NOTFOUND - Entity does not exist.
	NOTFOUND = "not_found"
	// UNKNOWN - Application unknown error.
	UNKNOWN = "unknown"
	// UNAUTHORIZED - User is not authorized.
	UNAUTHORIZED = "unauthorized"
	// FORBIDDEN - User is forbidden from performing an action.
	FORBIDDEN = "forbidden"
)

var (
	// DefaultCode is the default code returned when
	// none is specified.
	DefaultCode = INTERNAL
	// GlobalError is a general message when no error message
	// has been found.
	GlobalError = "An error has occurred."
)

// Error defines a standard application error.
type Error struct {
	// The application error code.
	Code string `json:"code" bson:"code"`
	// A human-readable message to send back to the end user.
	Message string `json:"message" bson:"message"`
	// Defines what operation is currently being run.
	Operation string `json:"operation" bson:"op"`
	// The error that was returned from the caller.
	Err      error `json:"error" bson:"error"`
	fileLine string
	pcs      []uintptr
}

// Error returns the string representation of the error
// message by implementing the error interface.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the error code if there is one.
	if e.Code != "" {
		buf.WriteString("<" + e.Code + "> ")
	}

	// Print the file-line, if any.
	if e.fileLine != "" {
		buf.WriteString(e.fileLine + " - ")
	}

	// Print the current operation in our stack, if any.
	if e.Operation != "" {
		buf.WriteString(e.Operation + ": ")
	}

	// Print the original error message, if any.
	if e.Err != nil {
		buf.WriteString(e.Err.Error() + ", ")
	}

	// Print the message, if any.
	if e.Message != "" {
		buf.WriteString(e.Message)
	}

	return strings.TrimSuffix(strings.TrimSpace(buf.String()), ",")
}

// HTTPStatusCode is a convenience method used to get the appropriate
// HTTP response status code for the respective error type.
func (e *Error) HTTPStatusCode() int {
	status := http.StatusInternalServerError
	switch e.Code {
	case CONFLICT:
		return http.StatusConflict
	case INVALID:
		return http.StatusBadRequest
	case NOTFOUND:
		return http.StatusNotFound
	}
	return status
}

// NewE returns an Error with the DefaultCode.
func NewE(err error, message, op string) *Error {
	return newError(err, message, DefaultCode, op)
}

// FileLine returns the file and line in which the error
// occurred.
func (e *Error) FileLine() string {
	return e.fileLine
}

// RuntimeFrames returns function/file/line information.
func (e *Error) RuntimeFrames() *runtime.Frames {
	return runtime.CallersFrames(e.pcs)
}

// ProgramCounters returns the slice of PC values associated
// with the error.
func (e *Error) ProgramCounters() []uintptr {
	return e.pcs
}

// StackTrace returns a string representation of the errors
// stacktrace, where each trace is separated by a newline
// and tab '\t'.
func (e *Error) StackTrace() string {
	trace := make([]string, 0, 100)
	rFrames := e.RuntimeFrames()
	frame, ok := rFrames.Next()
	line := strconv.Itoa(frame.Line)
	trace = append(trace, frame.Function+"(): "+e.Message)

	for ok {
		trace = append(trace, "\t"+frame.File+":"+line)
		frame, ok = rFrames.Next()
	}

	return strings.Join(trace, "\n")
}
