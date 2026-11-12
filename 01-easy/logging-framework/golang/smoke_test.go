package loggingframework

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()
	logger := NewLogger(nil)
	if err := logger.Info("test"); err != nil {
		t.Fatalf("log info: %v", err)
	}
}

func TestLogLevelFiltering(t *testing.T) {
	t.Parallel()
	appender := &recordingAppender{}
	logger := NewLogger(NewLoggerConfig(LogLevelWarning, appender))
	_ = logger.Info("skipped")
	_ = logger.Warning("kept")
	if len(appender.msgs) != 1 || !strings.Contains(appender.msgs[0], "kept") {
		t.Fatalf("msgs = %v, want one warning", appender.msgs)
	}
}

func TestFileAppender(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "test.log")
	appender := NewFileAppender(path)
	logger := NewLogger(NewLoggerConfig(LogLevelInfo, appender))
	if err := logger.Info("hello file"); err != nil {
		t.Fatalf("info: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if !strings.Contains(string(data), "hello file") {
		t.Fatalf("file content = %q", data)
	}
}

type recordingAppender struct {
	msgs []string
}

func (a *recordingAppender) Append(message *LogMessage) error {
	a.msgs = append(a.msgs, message.Message)
	return nil
}
