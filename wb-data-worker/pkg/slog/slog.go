package slog

import (
	"log/slog"
	"os"
)

var (
	String      = slog.String
	StringValue = slog.StringValue
)

func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}
