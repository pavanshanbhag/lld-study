package loggingframework

import "testing"

func TestNewLogger(t *testing.T) {
	t.Parallel()
	logger := NewLogger(nil)
	if err := logger.Info("test"); err != nil {
		t.Fatalf("log info: %v", err)
	}
}
