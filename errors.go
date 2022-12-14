package errors

import (
	"errors"
	"fmt"
	"sync"
)

type appError struct {
	msg   string
	frame []frame
	err   error
	once  sync.Once
}

func New(msg string) error {
	e := &appError{
		msg: msg,
	}
	e.once.Do(e.trace)
	return e
}

func (e *appError) trace() {
	pcs := callers()
	e.frame = newFrame(pcs)
}

// TODO:
func Wrap(err error, msg string) error {
	var app *appError
	if errors.As(err, &app) {
		return app.Unwrap()
	}
	return New(err.Error())
}

// TODO:
func (e *appError) Error() string {
	fmt.Println(e.frame)
	return e.msg
}

// TODO:
func (e *appError) Unwrap() error { return e.err }
