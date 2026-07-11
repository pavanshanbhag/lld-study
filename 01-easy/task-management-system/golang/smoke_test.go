package taskmanagementsystem

import (
	"testing"
	"time"
)

func TestNewTaskManager(t *testing.T) {
	t.Parallel()
	tm := NewTaskManager()
	user := NewUser("u1", "Alice", "alice@example.com")
	task := NewTask("1", "title", "desc", time.Now(), 1, user)
	tm.CreateTask(task)
	if len(tm.SearchTasks("title")) != 1 {
		t.Fatal("expected one task after create")
	}
}
