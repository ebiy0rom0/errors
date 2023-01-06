package errors

import (
	"fmt"
	"sync"
)

// wrapper for error added to stack trace only
type withStack struct {
	err  error
	st   stackTrace
	once sync.Once
}

// WithStack returns the error interface
// that wrapped other error and added stack trace.
// If err is nil returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	e := &withStack{
		err: err,
	}
	e.trace()

	return e
}

// trace obtains and saves a stack trace.
// Stack trace are obtained only once at runtime.
func (e *withStack) trace() {
	e.once.Do(func() {
		pcs := callers()
		e.st = newFrame(pcs)
	})
}

// Unwrap returns error interface being wrapped.
// It's used by erros.Is, errors.As and more.
func (e *withStack) Unwrap() error { return e.err }

// Errors returns error message.
// It's an implementation of the error interface.
// Return only the message of the wrapped error since there is no message in withStack.
func (e *withStack) Error() string { return e.err.Error() }

// Format is specify formatting rule for print.
// It's an implementation of the fmt.Formatter interface.
func (e *withStack) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprintf(f, "cause='%s'\nstackTrace:\n%s", e.Error(), e.st.output())
	case 's':
		fmt.Fprint(f, e.Error())
	}
}
