package errors

import (
	"fmt"
	"sync"
)

type wrapError struct {
	msg  string
	st   stackTrace
	once sync.Once
	err  error
}

func Wrap(err error, msg string) error {
	if err == nil {
		e := &fundamental{
			msg: msg,
		}
		e.trace()
		return e
	}

	e := &wrapError{
		msg: msg,
		err: err,
	}
	e.trace()
	return e
}

func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		e := &fundamental{
			msg: fmt.Sprintf(format, args...),
		}
		e.trace()
		return e
	}

	e := &wrapError{
		msg: fmt.Sprintf(format, args...),
		err: err,
	}
	e.trace()
	return e
}

func (e *wrapError) trace() {
	e.once.Do(func() {
		pcs := callers()
		e.st = newFrame(pcs)
	})
}

func (e *wrapError) Error() string { return e.msg + ": " + e.err.Error() }

func (e *wrapError) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprintf(f, "cause='%s'\nstackTrace:\n%s", e.Error(), e.origin().output())
	case 's':
		fmt.Fprint(f, e.Error())
	}
}

func (e *wrapError) origin() stackTrace {
	var origin = e.st
	if e, ok := e.err.(*wrapError); ok {
		return e.origin()
	} else if e, ok := e.err.(*fundamental); ok {
		return e.st
	}
	return origin
}
