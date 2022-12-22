package errors

import (
	"fmt"
	"sync"
)

// error wrapper structure
type wrapError struct {
	msg  string
	st   stackTrace
	once sync.Once
	err  error
}

// Wrap returns the error interface that wrapped args error
// and added stack trace.
// If err is nil returns unwrapped error(=fundamental).
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

// Wrapf returns the error interface that wrapped args error
// and added stack trace.
// Standard formatting can be used for error message.
// If err is nil returns unwrapped error(=fundamental).
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

// trace obtains and saves a stack trace.
// Stack trace are obtained only once at runtime.
func (e *wrapError) trace() {
	e.once.Do(func() {
		pcs := callers()
		e.st = newFrame(pcs)
	})
}

// Unwrap returns error interface being wrapped.
// It's used by erros.Is, errors.As and more.
func (e *wrapError) Unwrap() error { return e.err }

// Errors returns error message.
// It's an implementation of the error interface.
func (e *wrapError) Error() string { return e.msg + ": " + e.err.Error() }

// Format is specify formatting rule for print.
// It's an implementation of the fmt.Formatter interface.
func (e *wrapError) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprintf(f, "cause='%s'\nstackTrace:\n%v", e.Error(), e.origin())
	case 's':
		fmt.Fprint(f, e.Error())
	}
}

// origin returns the stack trace that obtained at oldest.
func (e *wrapError) origin() stackTrace {
	if w, ok := e.err.(*wrapError); ok {
		return w.origin()
	} else if f, ok := e.err.(*fundamental); ok {
		return f.st
	}
	return e.st
}
