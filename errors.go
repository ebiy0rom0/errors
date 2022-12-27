package errors

import (
	"errors"
	"fmt"
	"sync"
)

// basic error struct
type fundamental struct {
	msg  string
	st   stackTrace
	once sync.Once
}

// New returns the error interface that added stack trace.
// Specify any string for the error message.
func New(msg string) error {
	e := &fundamental{
		msg: msg,
	}
	e.trace()
	return e
}

// Errorf returns the error interface that added stack trace.
// Standard formatting can be used for error message.
func Errorf(format string, args ...any) error {
	e := &fundamental{
		msg: fmt.Sprintf(format, args...),
	}
	e.trace()
	return e
}

// trace obtains and saves a stack trace.
// Stack trace are obtained only once at runtime.
func (e *fundamental) trace() {
	e.once.Do(func() {
		pcs := callers()
		e.st = newFrame(pcs)
	})
}

// Errors returns error message.
// It's an implementation of the error interface.
func (e *fundamental) Error() string { return e.msg }

// Format is specify formatting rule for print.
// It's an implementation of the fmt.Formatter interface.
func (e *fundamental) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprintf(f, "cause='%s'\nstackTrace:\n%s", e.Error(), e.st.output())
	case 's':
		fmt.Fprint(f, e.Error())
	}
}

// Is the wrapper for the standard errors.Is().
// It's returns same result.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As the wrapper for the standard errors.As().
// It's returns same result.
func As(err error, target any) bool {
	return errors.As(err, target)
}
