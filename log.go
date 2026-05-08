package log

import "fmt"

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
