package cli

import (
	"cmp"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var appLogger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))

func configureLogger() error {
	levelText := cmp.Or(strings.TrimSpace(strings.ToLower(viper.GetString("log_level"))), "info")
	format := cmp.Or(strings.TrimSpace(strings.ToLower(viper.GetString("log_format"))), "text")

	level := new(slog.LevelVar)
	switch levelText {
	case "info":
		level.Set(slog.LevelInfo)
	case "debug":
		level.Set(slog.LevelDebug)
	case "warn", "warning":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	default:
		return fmt.Errorf("invalid log level: %q (supported: debug, info, warn, error)", levelText)
	}

	handler, err := buildLogHandler(format, level)
	if err != nil {
		return err
	}
	appLogger = slog.New(handler)
	return nil
}

func buildLogHandler(format string, level *slog.LevelVar) (slog.Handler, error) {
	opts := &slog.HandlerOptions{Level: level}
	switch format {
	case "text":
		return slog.NewTextHandler(os.Stderr, opts), nil
	case "json":
		return slog.NewJSONHandler(os.Stderr, opts), nil
	default:
		return nil, fmt.Errorf("invalid log format: %q (supported: text, json)", format)
	}
}

func logger() *slog.Logger {
	return appLogger
}

func logInfo(msg string, args ...any) {
	logger().Info(msg, args...)
}

func logWarn(msg string, args ...any) {
	logger().Warn(msg, args...)
}

func logError(msg string, args ...any) {
	logger().Error(msg, args...)
}

func logOutputWriter() io.Writer {
	return os.Stdout
}
