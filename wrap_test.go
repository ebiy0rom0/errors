package errors

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

// this library's error
var (
	Err1 = New("error1")
	Err2 = New("error2")
)

// standard error
var (
	ErrStd1 = errors.New("std error1")
	ErrStd2 = errors.New("std error2")
)

func TestWrap(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	type testcase []struct {
		name string
		args args
	}
	tests := testcase{
		{name: "wrap lib error", args: args{err: Err1, msg: "wrap"}},
		{name: "wrap multiple", args: args{err: Wrap(Err2, "first wrap"), msg: "second wrap"}},
		{name: "wrap stderror", args: args{err: ErrStd1, msg: "wrap std"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.err, tt.args.msg)
			if !errors.Is(err, tt.args.err) {
				t.Errorf("Unwrap() may not be implemented.")
			}

			msg := fmt.Sprintf("%s: %s", tt.args.msg, tt.args.err.Error())
			if err.Error() != msg {
				t.Errorf("unmatch Error() message. out=%s, want=%s", err.Error(), msg)
			}

			// check implements fmt.Formatter
			var buf bytes.Buffer
			if _, err := buf.WriteString(fmt.Sprintf("%v", err)); err != nil {
				t.Errorf("error unexpected. err=%v", err)
			}
			if _, err := buf.WriteString(fmt.Sprintf("%s", err)); err != nil {
				t.Errorf("error unexpected. err=%s", err)
			}

			// output log
			t.Log(buf.String())
		})
	}

	tests = testcase{
		{name: "nil wrap", args: args{err: nil, msg: "nil wrap"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.err, tt.args.msg)

			// if the argument err is nil, the fundamental is returned.
			if err.Error() != tt.args.msg {
				t.Errorf("unmatch Error() message. out=%s, want=%s", err.Error(), tt.args.msg)
			}

			// check implements fmt.Formatter
			var buf bytes.Buffer
			if _, err := buf.WriteString(fmt.Sprintf("%v", err)); err != nil {
				t.Errorf("error unexpected. err=%v", err)
			}
			if _, err := buf.WriteString(fmt.Sprintf("%s", err)); err != nil {
				t.Errorf("error unexpected. err=%s", err)
			}

			// output log
			t.Log(buf.String())
		})
	}
}

func TestWrapf(t *testing.T) {
	type args struct {
		err    error
		format string
		args   []any
	}
	type testcase []struct {
		name string
		args args
	}
	tests := testcase{
		{name: "string format", args: args{err: Err1, format: "%s", args: []any{"wrap"}}},
		{name: "number format", args: args{err: Err2, format: "%d", args: []any{123}}},
		{name: "struct format", args: args{err: Err1, format: "%v", args: []any{struct{ msg string }{msg: "wrap"}}}},
		{name: "multiple format", args: args{err: Err1, format: "%s %d %v", args: []any{"wrap", 123, struct{ msg string }{msg: "wrap"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrapf(tt.args.err, tt.args.format, tt.args.args...)
			if !errors.Is(err, tt.args.err) {
				t.Errorf("Unwrap() may not be implemented.")
			}

			msg := fmt.Sprintf("%s: %s", fmt.Sprintf(tt.args.format, tt.args.args...), tt.args.err.Error())
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
		{name: "nil wrap format", args: args{err: nil, format: "%s", args: []any{"nil wrap"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrapf(tt.args.err, tt.args.format, tt.args.args...)

			// if the argument err is nil, the fundamental is returned.
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
			t.Logf("%v", buf.String())
		})
	}
}
