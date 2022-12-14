package errors

import (
	"runtime"
)

type frame struct {
	pc   uintptr
	name string
	file string
	line int
}

func callers() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(0, pcs[:])
	return pcs[0 : n-2]
}

func newFrame(pcs []uintptr) (frames []frame) {
	for _, pc := range pcs {
		f := frame{pc: pc}
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			return frames
		}

		f.file, f.line = fn.FileLine(pc)
		f.name = fn.Name()

		frames = append(frames, f)
	}
	return
}
