package loggingframework

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	appLogger := NewLogger(nil)

	_ = appLogger.Info("This is an information message")
	_ = appLogger.Warning("This is a warning message")
	_ = appLogger.Error("This is an error message")

	fileAppender := NewFileAppender("app.log")
	config := NewLoggerConfig(LogLevelDebug, fileAppender)
	appLogger.SetConfig(config)

	if err := appLogger.Debug("This is a debug message"); err != nil {
		logger.Error("debug log failed", "err", err)
	}
	if err := appLogger.Info("This is an information message"); err != nil {
		logger.Error("info log failed", "err", err)
	}
}
