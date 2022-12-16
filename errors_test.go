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
		{name: "new error", msg: "error message"},
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

			var buf bytes.Buffer
			if _, err := buf.WriteString(fmt.Sprintf("%v", err)); err != nil {
				t.Errorf("error unexpected. err=%v", err)
			} else {
				t.Logf("%v", buf.String())
			}
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		args   []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Errorf(tt.args.format, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Errorf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
