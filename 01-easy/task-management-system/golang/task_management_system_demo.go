package taskmanagementsystem

import (
	"log/slog"
	"os"
	"time"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	taskManager := NewTaskManager()

	user1 := NewUser("1", "John Doe", "john@example.com")
	user2 := NewUser("2", "Jane Smith", "jane@example.com")

	task1 := NewTask("1", "Task 1", "Description 1", time.Now().Add(24*time.Hour), 1, user1)
	task2 := NewTask("2", "Task 2", "Description 2", time.Now().Add(48*time.Hour), 2, user2)
	task3 := NewTask("3", "Task 3", "Description 3", time.Now().Add(72*time.Hour), 1, user1)

	_ = taskManager.CreateTask(task1)
	_ = taskManager.CreateTask(task2)
	_ = taskManager.CreateTask(task3)

	task2.Description = "Updated description"
	if err := taskManager.UpdateTask(task2); err != nil {
		logger.Error("update failed", "err", err)
	}

	for _, task := range taskManager.SearchTasks("Task") {
		logger.Info("search hit", "title", task.Title)
	}

	for _, task := range taskManager.FilterTasks(Pending, time.Unix(0, 0), time.Now().Add(96*time.Hour), 1) {
		logger.Info("filter hit", "title", task.Title)
	}

	if err := taskManager.MarkTaskAsCompleted("1"); err != nil {
		logger.Error("complete failed", "err", err)
	}

	for _, task := range taskManager.TaskHistory(user1) {
		logger.Info("history", "user", user1.Name, "title", task.Title, "status", task.Status)
	}

	if err := taskManager.DeleteTask("3"); err != nil {
		logger.Error("delete failed", "err", err)
	}
}
