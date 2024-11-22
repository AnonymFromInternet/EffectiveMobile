package logger

import "log/slog"

func WrapError(e error) slog.Attr {
	return slog.Attr{Key: "error", Value: slog.StringValue(e.Error())}
}
