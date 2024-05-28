package stp

import (
	"runtime"
	"strconv"
	"strings"
)

func GoroutineID() (uint64, error) {
	// Buffer to hold stack trace
	buf := make([]byte, 64)
	// Read the current goroutine stack trace
	buf = buf[:runtime.Stack(buf, false)]
	// Convert stack trace to string and find the goroutine ID
	idField := strings.Fields(strings.TrimPrefix(string(buf), "goroutine "))[0]
	id, err := strconv.ParseUint(idField, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
