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

func (e *wrapError) Unwrap() error { return e.err }
func (e *wrapError) Error() string { return e.msg + ": " + e.err.Error() }

func (e *wrapError) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprintf(f, "cause='%s'\nstackTrace:\n%v", e.Error(), e.origin())
	case 's':
		fmt.Fprint(f, e.Error())
	}
}

func (e *wrapError) origin() stackTrace {
	if w, ok := e.err.(*wrapError); ok {
		return w.origin()
	} else if f, ok := e.err.(*fundamental); ok {
		return f.st
	}
	return e.st
}
