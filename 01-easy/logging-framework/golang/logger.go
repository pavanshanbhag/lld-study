package loggingframework

import (
	"sync"
)

type Logger struct {
	config *LoggerConfig
	mu     sync.RWMutex
}

func NewLogger(config *LoggerConfig) *Logger {
	if config == nil {
		config = NewLoggerConfig(LogLevelInfo, NewConsoleAppender())
	}
	return &Logger{config: config}
}

func (l *Logger) SetConfig(config *LoggerConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config = config
}

func (l *Logger) log(level LogLevel, message string) error {
	l.mu.RLock()
	if level < l.config.Level {
		l.mu.RUnlock()
		return nil
	}
	appender := l.config.Appender
	l.mu.RUnlock()

	logMessage := NewLogMessage(level, message)
	return appender.Append(logMessage)
}

func (l *Logger) Debug(message string) error {
	return l.log(LogLevelDebug, message)
}

func (l *Logger) Info(message string) error {
	return l.log(LogLevelInfo, message)
}

func (l *Logger) Warning(message string) error {
	return l.log(LogLevelWarning, message)
}

func (l *Logger) Error(message string) error {
	return l.log(LogLevelError, message)
}

func (l *Logger) Fatal(message string) error {
	return l.log(LogLevelFatal, message)
}
