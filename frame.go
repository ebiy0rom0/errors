package errors

import (
	"fmt"
	"runtime"
)

type frame struct {
	pc   uintptr
	name string
	file string
	line int
}

type stackTrace []frame

func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(7, pcs[:])
	return pcs[0 : n-2]
}

func newFrame(pcs []uintptr) (st stackTrace) {
	for _, pc := range pcs {
		f := frame{pc: pc}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return
		}

		f.file, f.line = fn.FileLine(pc)
		f.name = fn.Name()

		st = append(st, f)
	}
	return
}

func (f frame) output(no int) string {
	return fmt.Sprintf("#%02d %s, line:%d", no, f.file, f.line)
}

func (st stackTrace) output() string {
	var text = make([]byte, 0, 512)
	for no, f := range st {
		text = append(text, f.output(no+1)...)
		text = append(text, '\n')
	}
	return string(text)
}

func (st stackTrace) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprint(f, st.output())
	}
}
