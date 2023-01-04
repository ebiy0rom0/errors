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
		{name: "std error", args: args{err: stdErr, target: fmt.Errorf("wrap: %w", stdErr)}},
		{name: "predefined error", args: args{err: os.ErrNotExist, target: fmt.Errorf("wrap: %w", os.ErrNotExist)}},
		{name: "lib error", args: args{err: libErr, target: fmt.Errorf("wrap: %w", libErr)}},
		{name: "wrap std error", args: args{err: stdErr, target: Wrap(stdErr, "wrap std error")}},
		{name: "wrap predefined error", args: args{err: os.ErrNotExist, target: Wrap(os.ErrNotExist, "wrap predefined error")}},
		{name: "wrap lib error", args: args{err: libErr, target: Wrap(libErr, "wrap lib error")}},
		{name: "wrapf std error", args: args{err: stdErr, target: Wrapf(stdErr, "%s", "wrapf std error")}},
		{name: "wrapf predefined error", args: args{err: os.ErrNotExist, target: Wrapf(os.ErrNotExist, "%s", "wrapf predefined error")}},
		{name: "wrapf lib error", args: args{err: libErr, target: Wrapf(libErr, "%s", "wrapf lib error")}},
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
	err := New("lib error")
	type args struct {
		err    error
		target any
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "lib error", args: args{err: err, target: new(fundamental)}},
		{name: "stdwrap error", args: args{err: fmt.Errorf("wrap: %w", err), target: new(fundamental)}},
		{name: "wrap lib error", args: args{err: Wrap(err, "wrap error"), target: new(fundamental)}},
		{name: "wrapf lib error", args: args{err: Wrapf(err, "%s", "wrapf error"), target: new(fundamental)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := errors.As(tt.args.err, &tt.args.target)
			if got := As(tt.args.err, &tt.args.target); got != want {
				t.Errorf("As() returns different results of errors.As(). result=%v, want %v", got, want)
			}
		})
	}
}
