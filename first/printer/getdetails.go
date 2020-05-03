package printer

import (
	"runtime"
)

// GoVersion prints the current version of go
// covering the printer exrcise
func GoVersion() string {
	return runtime.Version()
}
