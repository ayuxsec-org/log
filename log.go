package log

import (
	"fmt"
	"os"
)

func Info(msg string) {
	if DisableInfo {
		return
	}
	fmt.Fprint(Target, infoPrefix+format(msg))
}

func Infof(formatStr string, a ...any) {
	if DisableInfo {
		return
	}
	fmt.Fprint(Target, infoPrefix+format(fmt.Sprintf(formatStr, a...)))
}

func Warn(msg string) {
	if DisableWarn {
		return
	}
	fmt.Fprint(Target, warnPrefix+format(msg))
}

func Warnf(formatStr string, a ...any) {
	if DisableWarn {
		return
	}
	fmt.Fprint(Target, warnPrefix+format(fmt.Sprintf(formatStr, a...)))
}

func Debug(msg string) {
	if DisableDebug {
		return
	}
	fmt.Fprint(Target, debugPrefix+format(msg))
}

func Debugf(formatStr string, a ...any) {
	if DisableDebug {
		return
	}
	fmt.Fprint(Target, debugPrefix+format(fmt.Sprintf(formatStr, a...)))
}

func Error(msg string) {
	if DisableError {
		return
	}
	fmt.Fprint(Target, errorPrefix+format(msg))
}

func Errorf(formatStr string, a ...any) {
	if DisableError {
		return
	}
	fmt.Fprint(Target, errorPrefix+format(fmt.Sprintf(formatStr, a...)))
}

func Fatal(msg string) {
	fmt.Fprint(Target, fatalPrefix+format(msg))
	os.Exit(1)
}

func Fatalf(formatStr string, a ...any) {
	fmt.Fprint(Target, fatalPrefix+format(fmt.Sprintf(formatStr, a...)))
	os.Exit(1)
}
