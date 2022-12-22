package errors

import (
	"fmt"
	"runtime"
)

// structure of the frame to be obtained
type frame struct {
	pc   uintptr
	name string
	file string
	line int
}

// structure of stack trace
type stackTrace []frame

// caller returns program counters from the point where the function was called.
// skip values are customized for this package, don't use them elsewhere and
// keep the calling hierarchy consistent.
func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(7, pcs[:])
	return pcs[0 : n-2]
}

// newFrame returns the stack trace.
func newFrame(pcs []uintptr) (st stackTrace) {
	for _, pc := range pcs {
		f := frame{pc: pc - 1}
		fn := runtime.FuncForPC(f.pc)
		if fn == nil {
			return
		}

		f.file, f.line = fn.FileLine(f.pc)
		f.name = fn.Name()

		st = append(st, f)
	}
	return
}

// output returns the formed stack trace.
func (f frame) output(no int) string {
	return fmt.Sprintf("\t#%02d %s(%d): %s", no, f.file, f.line, f.name)
}

// output returns the formed stack trace for all.
func (st stackTrace) output() string {
	var text = make([]byte, 0, 512)
	for no, f := range st {
		text = append(text, f.output(no+1)...)
		text = append(text, '\n')
	}
	return string(text)
}

// Format is specify formatting rule for print.
// It's an implementation of the fmt.Formatter interface.
func (st stackTrace) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprint(f, st.output())
	}
}
