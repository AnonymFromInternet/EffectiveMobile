package logger

import "log/slog"

func WrapError(e error) slog.Attr {
	return slog.Attr{Key: "error", Value: slog.StringValue(e.Error())}
}

func WrapDebug(msg string) slog.Attr {
	return slog.Attr{Key: "debug", Value: slog.StringValue(msg)}
}
