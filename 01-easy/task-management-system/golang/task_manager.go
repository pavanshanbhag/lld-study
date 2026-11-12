package taskmanagementsystem

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type TaskManager struct {
	tasks     map[string]*Task
	userTasks map[string][]*Task
	mu        sync.Mutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:     make(map[string]*Task),
		userTasks: make(map[string][]*Task),
	}
}

func (tm *TaskManager) CreateTask(task *Task) error {
	if task == nil || task.AssignedUser == nil {
		return fmt.Errorf("task and assigned user required")
	}
	tm.mu.Lock()
	defer tm.mu.Unlock()

	tm.tasks[task.Id] = task
	tm.assignTaskToUser(task.AssignedUser, task)
	return nil
}

func (tm *TaskManager) UpdateTask(updatedTask *Task) error {
	if updatedTask == nil {
		return fmt.Errorf("task required")
	}
	tm.mu.Lock()
	defer tm.mu.Unlock()

	existingTask, exists := tm.tasks[updatedTask.Id]
	if !exists {
		return fmt.Errorf("task not found: %s", updatedTask.Id)
	}

	existingTask.Title = updatedTask.Title
	existingTask.Description = updatedTask.Description
	existingTask.DueDate = updatedTask.DueDate
	existingTask.Priority = updatedTask.Priority
	existingTask.Status = updatedTask.Status

	if existingTask.AssignedUser != updatedTask.AssignedUser {
		tm.unassignTaskFromUser(existingTask.AssignedUser, existingTask)
		tm.assignTaskToUser(updatedTask.AssignedUser, existingTask)
		existingTask.AssignedUser = updatedTask.AssignedUser
	}
	return nil
}

func (tm *TaskManager) DeleteTask(taskId string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[taskId]
	if !exists {
		return fmt.Errorf("task not found: %s", taskId)
	}
	delete(tm.tasks, taskId)
	tm.unassignTaskFromUser(task.AssignedUser, task)
	return nil
}

func (tm *TaskManager) SearchTasks(keyword string) []*Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	var matchingTasks []*Task
	for _, task := range tm.tasks {
		if strings.Contains(task.Title, keyword) || strings.Contains(task.Description, keyword) {
			matchingTasks = append(matchingTasks, task)
		}
	}
	return matchingTasks
}

func (tm *TaskManager) FilterTasks(status TaskStatus, startDate, endDate time.Time, priority int) []*Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	var filteredTasks []*Task
	for _, task := range tm.tasks {
		if task.Status == status &&
			task.DueDate.After(startDate) &&
			task.DueDate.Before(endDate) &&
			task.Priority == priority {
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}

func (tm *TaskManager) MarkTaskAsCompleted(taskId string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[taskId]
	if !exists {
		return fmt.Errorf("task not found: %s", taskId)
	}
	task.Status = Completed
	return nil
}

func (tm *TaskManager) TaskHistory(user *User) []*Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	return tm.userTasks[user.Id]
}

func (tm *TaskManager) assignTaskToUser(user *User, task *Task) {
	tm.userTasks[user.Id] = append(tm.userTasks[user.Id], task)
}

func (tm *TaskManager) unassignTaskFromUser(user *User, task *Task) {
	tasks := tm.userTasks[user.Id]
	for i, t := range tasks {
		if t == task {
			tm.userTasks[user.Id] = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
}
