package errors

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestError_Error(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed: %s", err.Error())
	}

	tt := map[string]struct {
		input *Error
		want  string
	}{
		"Normal": {
			NewInternal(fmt.Errorf("error"), "message", "op"),
			"<internal> " + wd + "/errors_test.go:21 - op: error, message",
		},
		"Nil Operation": {
			NewInternal(fmt.Errorf("error"), "message", ""),
			"<internal> " + wd + "/errors_test.go:25 - error, message",
		},
		"Nil Err": {
			NewInternal(nil, "message", ""),
			"<internal> " + wd + "/errors_test.go:29 - message",
		},
		"Nil Message": {
			NewInternal(fmt.Errorf("error"), "", ""),
			"<internal> " + wd + "/errors_test.go:33 - error",
		},
		"Message Error": {
			&Error{Message: "message", Err: fmt.Errorf("err")},
			"err, message",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.Error()
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expecting %s, got %s", test.want, got)
			}
		})
	}
}

func TestError_FileLine(t *testing.T) {
	e := &Error{fileLine: "fileline:20"}
	got := e.FileLine()
	want := "fileline:20"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expecting %s, got %s", want, got)
	}
}

func TestError_HTTPStatusCode(t *testing.T) {
	tt := map[string]struct {
		input Error
		want  int
	}{
		"Conflict": {
			Error{Code: CONFLICT},
			http.StatusConflict,
		},
		"Internal": {
			Error{Code: INTERNAL},
			http.StatusInternalServerError,
		},
		"Invalid": {
			Error{Code: INVALID},
			http.StatusBadRequest,
		},
		"Not Found": {
			Error{Code: NOTFOUND},
			http.StatusNotFound,
		},
		"Unknown": {
			Error{Code: UNKNOWN},
			http.StatusInternalServerError,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := test.input.HTTPStatusCode()
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expecting %d, got %d", test.want, got)
			}
		})
	}
}

func TestError_ProgramCounters(t *testing.T) {
	e := NewE(fmt.Errorf("error"), "message", "op")
	got := e.ProgramCounters()
	want := 100
	if !reflect.DeepEqual(len(got), want) {
		t.Fatalf("expecting %d, got %d", want, got)
	}
}

func TestError_RuntimeFrames(t *testing.T) {
	e := NewE(fmt.Errorf("error"), "message", "op")
	got := e.RuntimeFrames()
	frame, _ := got.Next()
	want := "github.com/terrapi-solution/controller/domain/errors.TestError_RuntimeFrames"
	if !reflect.DeepEqual(want, frame.Function) {
		t.Fatalf("expecting %s, got %s", want, frame.Function)
	}
}

func TestError_StackTrace(t *testing.T) {
	e := NewE(fmt.Errorf("error"), "message", "op")
	got := e.StackTrace()
	want := "github.com/terrapi-solution/controller/domain/errors.TestError_StackTrace(): message"
	if !strings.Contains(got, want) {
		t.Fatalf("expecting %s to contain, got %s", want, got)
	}
}
