package log

import "os"

var Target = os.Stderr

var (
	WithTimeStamp = false
	WithTrace     = true
)

var (
	DisableInfo  = false
	DisableDebug = false
	DisableWarn  = false
	DisableError = false
)
