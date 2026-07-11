package courseregistrationsystem

import "testing"

func TestCourseRegistrationSystemRegisterCourse(t *testing.T) {
	t.Parallel()

	system := NewCourseRegistrationSystem()
	course := NewCourse("CS101", "Intro to Programming", "Dr. Smith", 30)
	student := NewStudent(1, "Alice", "alice@example.com")

	system.AddCourse(course)
	system.AddStudent(student)

	if err := system.RegisterCourse(student, course); err != nil {
		t.Fatalf("RegisterCourse: %v", err)
	}

	registered := system.GetRegisteredCourses(student)
	if len(registered) != 1 || registered[0].Code != "CS101" {
		t.Fatalf("GetRegisteredCourses() = %+v, want CS101", registered)
	}
}
