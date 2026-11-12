package taskmanagementsystem

import (
	"testing"
	"time"
)

func TestNewTaskManager(t *testing.T) {
	t.Parallel()
	tm := NewTaskManager()
	user := NewUser("u1", "Alice", "alice@example.com")
	task := NewTask("1", "title", "desc", time.Now().Add(time.Hour), 1, user)
	if err := tm.CreateTask(task); err != nil {
		t.Fatalf("create: %v", err)
	}
	if len(tm.SearchTasks("title")) != 1 {
		t.Fatal("expected one task after create")
	}
}

func TestMarkCompletedAndHistory(t *testing.T) {
	t.Parallel()
	tm := NewTaskManager()
	user := NewUser("u1", "Alice", "alice@example.com")
	task := NewTask("1", "Ship", "ship it", time.Now().Add(time.Hour), 1, user)
	_ = tm.CreateTask(task)
	if err := tm.MarkTaskAsCompleted("1"); err != nil {
		t.Fatalf("complete: %v", err)
	}
	if task.Status != Completed {
		t.Fatalf("status = %v, want Completed", task.Status)
	}
	history := tm.TaskHistory(user)
	if len(history) != 1 || history[0].Id != "1" {
		t.Fatal("expected task in history")
	}
}

func TestDeleteMissingTask(t *testing.T) {
	t.Parallel()
	tm := NewTaskManager()
	if err := tm.DeleteTask("missing"); err == nil {
		t.Fatal("expected error for missing task")
	}
}
