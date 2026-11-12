package courseregistrationsystem

import (
	"fmt"
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	registrationSystem := NewCourseRegistrationSystem()

	// Create courses
	course1 := NewCourse("CS101", "Introduction to Programming", "John Doe", 50)
	course2 := NewCourse("CS201", "Data Structures and Algorithms", "Jane Smith", 30)

	registrationSystem.AddCourse(course1)
	registrationSystem.AddCourse(course2)

	// Create students
	student1 := NewStudent(1, "Alice", "alice@example.com")
	student2 := NewStudent(2, "Bob", "bob@example.com")

	registrationSystem.AddStudent(student1)
	registrationSystem.AddStudent(student2)

	// Search for courses
	searchResults := registrationSystem.SearchCourses("CS")
	logger.Info("demo", "msg", fmt.Sprint("Search Results:"))
	for _, course := range searchResults {
		logger.Info("demo", "msg", fmt.Sprintf("%s - %s\n", course.Code, course.Name))
	}

	// Register courses for students
	logger.Info("demo", "msg", fmt.Sprint("\nRegistration Results:"))

	err := registrationSystem.RegisterCourse(student1, course1)
	logger.Info("demo", "msg", fmt.Sprintf("Student 1 - Course 1: %v\n", err == nil))

	err = registrationSystem.RegisterCourse(student2, course1)
	logger.Info("demo", "msg", fmt.Sprintf("Student 2 - Course 1: %v\n", err == nil))

	err = registrationSystem.RegisterCourse(student1, course2)
	logger.Info("demo", "msg", fmt.Sprintf("Student 1 - Course 2: %v\n", err == nil))

	// Get registered courses for a student
	registeredCourses := registrationSystem.GetRegisteredCourses(student1)
	logger.Info("demo", "msg", fmt.Sprint("\nRegistered Courses for Student 1:"))
	for _, course := range registeredCourses {
		logger.Info("demo", "msg", fmt.Sprintf("%s - %s\n", course.Code, course.Name))
	}
}
