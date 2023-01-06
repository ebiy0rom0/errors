package errors

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestWithStack(t *testing.T) {
	stdErr := errors.New("error")
	libErr := New("lib error")

	type testcase []struct {
		name string
		err  error
	}
	tests := testcase{
		{name: "std error", err: stdErr},
		{name: "lib error", err: libErr},
		{name: "error wrapping std", err: Wrap(stdErr, "std wrap")},
		{name: "error wrapping lib", err: Wrap(libErr, "lib wrap")},
		{name: "error wrapping fmt", err: fmt.Errorf("wrap is:%w", stdErr)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithStack(tt.err)
			if !errors.Is(err, tt.err) {
				t.Error("Unwrap() may not be implemented.")
			}

			msg := tt.err.Error()
			if err.Error() != msg {
				t.Errorf("unmatch Error() message. out=%s, want=%s", err.Error(), msg)
			}

			// check implements fmt.Formatter
			var buf bytes.Buffer
			if _, err := buf.WriteString(fmt.Sprintf("%s\n", err)); err != nil {
				t.Errorf("error unexpected. err=%s", err)
			}
			if _, err := buf.WriteString(fmt.Sprintf("%v", err)); err != nil {
				t.Errorf("error unexpected. err=%v", err)
			}

			// output log
			t.Logf("%v", buf.String())
		})
	}

	tests = testcase{
		{name: "args error is nil", err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithStack(tt.err)
			if err != nil {
				t.Errorf("WithStack returns not nil:%v", err)
			}
		})
	}
}
