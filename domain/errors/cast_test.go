package errors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestError_ToError(t *testing.T) {
	tt := map[string]struct {
		input any
		want  *Error
	}{
		"Pointer": {
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
		},
		"Non Pointer": {
			Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
			&Error{Code: INTERNAL, Message: "test", Operation: "op", Err: fmt.Errorf("err")},
		},
		"Default": {
			nil,
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ToError(test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expecting %s, got %s", test.want, got)
			}
		})
	}
}
