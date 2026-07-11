package stackoverflow

import "testing"

func TestNewStackOverflow(t *testing.T) {
	t.Parallel()
	so := NewStackOverflow()
	user := so.CreateUser("alice", "alice@example.com")
	if user == nil {
		t.Fatal("expected user")
	}
}
