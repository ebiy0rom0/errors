package errors

import (
	"fmt"
	"sync"
)

type fundamental struct {
	msg  string
	st   stackTrace
	once sync.Once
}

func New(msg string) error {
	e := &fundamental{
		msg: msg,
	}
	e.trace()
	return e
}

func Errorf(format string, args ...any) error {
	e := &fundamental{
		msg: fmt.Sprintf(format, args...),
	}
	e.trace()
	return e
}

func (e *fundamental) trace() {
	e.once.Do(func() {
		pcs := callers()
		e.st = newFrame(pcs)
	})
}

func (e *fundamental) Error() string { return e.msg }
