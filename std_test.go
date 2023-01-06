package errors

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

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
		{name: "std error", args: args{err: fmt.Errorf("wrap: %w", stdErr), target: stdErr}},
		{name: "predefined error", args: args{err: fmt.Errorf("wrap: %w", os.ErrNotExist), target: os.ErrNotExist}},
		{name: "lib error", args: args{err: fmt.Errorf("wrap: %w", libErr), target: libErr}},
		{name: "wrap std error", args: args{err: Wrap(stdErr, "wrap std error"), target: stdErr}},
		{name: "wrap predefined error", args: args{err: Wrap(os.ErrNotExist, "wrap predefined error"), target: os.ErrNotExist}},
		{name: "wrap lib error", args: args{err: Wrap(libErr, "wrap lib error"), target: libErr}},
		{name: "wrapf std error", args: args{err: Wrapf(stdErr, "%s", "wrapf std error"), target: stdErr}},
		{name: "wrapf predefined error", args: args{err: Wrapf(os.ErrNotExist, "%s", "wrapf predefined error"), target: os.ErrNotExist}},
		{name: "wrapf lib error", args: args{err: Wrapf(libErr, "%s", "wrapf lib error"), target: libErr}},
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

type customError struct {
	msg string
}

func (c customError) Error() string { return c.msg }

func TestAs(t *testing.T) {
	err := customError{msg: "error"}
	type args struct {
		err    error
		target any
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "lib error", args: args{err: err, target: new(customError)}},
		{name: "stdwrap error", args: args{err: fmt.Errorf("wrap: %w", err), target: new(customError)}},
		{name: "wrap lib error", args: args{err: Wrap(err, "wrap error"), target: new(customError)}},
		{name: "wrapf lib error", args: args{err: Wrapf(err, "%s", "wrapf error"), target: new(customError)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := errors.As(tt.args.err, tt.args.target)
			if got := As(tt.args.err, tt.args.target); got != want {
				t.Errorf("As() returns different results of errors.As(). result=%v, want %v", got, want)
			}

			f := tt.args.target.(*customError)
			if !reflect.DeepEqual(err, *f) {
				t.Errorf("set target error failed: %v", f)
			}
		})
	}
}
