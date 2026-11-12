package taskmanagementsystem

import "time"

type Task struct {
	Id           string
	Title        string
	Description  string
	DueDate      time.Time
	Priority     int
	Status       TaskStatus
	AssignedUser *User
}

func NewTask(id, title, description string, dueDate time.Time, priority int, assignedUser *User) *Task {
	return &Task{
		Id:           id,
		Title:        title,
		Description:  description,
		DueDate:      dueDate,
		Priority:     priority,
		Status:       Pending,
		AssignedUser: assignedUser,
	}
}
