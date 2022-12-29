package errors

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{name: "new error", msg: "error"},
		{name: "not found", msg: "file not found"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.msg)
			if err.Error() != tt.msg {
				t.Errorf("unmatch Error() message. out=%s, want=%s", err.Error(), tt.msg)
			} else {
				t.Log(err.Error())
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
			t.Log(buf.String())
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "string format", args: args{format: "%s", args: []any{"error"}}},
		{name: "number format", args: args{format: "%d", args: []any{123}}},
		{name: "struct format", args: args{format: "%v", args: []any{struct{ msg string }{msg: "error"}}}},
		{name: "multiple format", args: args{format: "%s %d %v", args: []any{"error", 123, struct{ msg string }{msg: "error"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Errorf(tt.args.format, tt.args.args...)
			msg := fmt.Sprintf(tt.args.format, tt.args.args...)
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
			t.Log(buf.String())
		})
	}
}

func TestIs(t *testing.T) {
	stdErr := errors.New("error")
	libErr := New("lib error")
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "std error", args: args{err: stdErr, target: stdErr}},
		{name: "predefined error", args: args{err: os.ErrNotExist, target: os.ErrNotExist}},
		{name: "lib error", args: args{err: libErr, target: libErr}},
		{name: "generated from predefined error", args: args{err: stdErr, target: Wrap(stdErr, "wrap error")}},
		{name: "error is nil", args: args{err: nil, target: nil}},
	}
	for _, tt := range tests {
		// Test cases are assumed to return true for all patterns.
		// If returned false, the error creator is likely to have a problem
		// and will be notified.
		t.Run(tt.name, func(t *testing.T) {
			want := errors.Is(tt.args.err, tt.args.target)
			if got := Is(tt.args.err, tt.args.target); got != want {
				t.Errorf("Is() returns different results of errors.Is(). result=%v, want %v", got, want)
			} else if !got {
				t.Log("Is() returns false.")
			}
		})
	}
}

func TestAs(t *testing.T) {
	stdErr := errors.New("error")
	libErr := New("lib error")
	type args struct {
		err    error
		target any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "std error", args: args{err: stdErr, target: stdErr}, want: true},
		{name: "predefined error", args: args{err: stdErr, target: os.ErrNotExist}, want: true},
		{name: "lib error", args: args{err: libErr, target: libErr}, want: true},
		{name: "defferent type error", args: args{err: stdErr, target: libErr}, want: false},
		{name: "std error", args: args{err: stdErr, target: stdErr}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := As(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("As() = %v, want %v", got, tt.want)
			}
		})
	}
}
