package errors

import (
	"runtime"
	"strings"
	"sync"
)

var pool sync.Pool

func inPackage(functionName, pkg string) bool {
	if !strings.HasPrefix(functionName, pkg) {
		return false
	}

	if len(functionName) == len(pkg) {
		return true
	}

	if len(functionName) > len(pkg) {
		if functionName[len(pkg)] == '.' || functionName[len(pkg)] == '/' {
			return true
		}
	}

	return false
}

func Locate() (locations []Location) {
	dataLen := 150
	temp1 := pool.Get()
	var data []uintptr
	if temp1 != nil {
		temp2 := temp1.([]uintptr)
		if len(temp2) >= dataLen {
			data = temp2
		}
	}
	if data == nil {
		data = make([]uintptr, dataLen)
	}

	pc := data[:dataLen]
	pc = pc[:runtime.Callers(3, pc)]

	frames := runtime.CallersFrames(pc)
	for {
		frame, more := frames.Next()

		if frame.Line != 0 && frame.File != "" && !inPackage(frame.Function, "runtime") {
			locations = append(locations, Location{
				file:     frame.File,
				line:     frame.Line,
				function: frame.Function,
			})
		}

		if len(locations) == dataLen {
			break
		}

		if !more {
			break
		}
	}

	pool.Put(data)

	return
}
