package batch



import (
	"github.com/unchartedsoftware/plog"
)

// logging utilities for the batch package

const (
	// Teal "BATCH" log prefix
	// Code derived from github.com/mgutz/ansi, but I wanted it as a const, so
	// couldn't se that directly
	preLog = "\033[1;38;5;6mBATCH\033[0m: "
)

func batchErrorf (format string, args ...interface{}) {
	log.Errorf(preLog+format+"\n", args...)
}
func batchWarnf (format string, args ...interface{}) {
	log.Warnf(preLog+format+"\n", args...)
}
func batchInfof (format string, args ...interface{}) {
	log.Infof(preLog+format+"\n", args...)
}
func batchDebugf (format string, args ...interface{}) {
	log.Debugf(preLog+format+"\n", args...)
}