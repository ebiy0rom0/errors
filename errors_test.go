package errors

import (
	"bytes"
	"fmt"
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
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAs(t *testing.T) {
	type args struct {
		err    error
		target any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := As(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("As() = %v, want %v", got, tt.want)
			}
		})
	}
}
