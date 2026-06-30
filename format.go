package log

import (
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

func format(msg string) string {
	var prefix string
	if WithTimeStamp {
		currentTimeStamp := time.Now().Format("2006-01-02 15:04:05")
		prefix = " " + currentTimeStamp
	}
	if WithTrace {
		file, line := trace()
		prefix += " " + file + ":" + strconv.Itoa(line) + ":"
	}
	return prefix + " " + msg + "\n"
}

func trace() (file string, line int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		panic("failed to trace log couldn't get caller info")
	}

	if !DisableDebug {
		return file, line
	}

	return filepath.Base(file), line
}
