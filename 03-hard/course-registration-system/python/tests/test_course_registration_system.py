from course import Course
from student import Student
from course_registration_system import CourseRegistrationSystem


def test_course_registration_system() -> None:
    system = CourseRegistrationSystem()
    course = Course("CS101", "Intro to Programming", "Dr. Smith", 30, 0)
    student = Student(1, "Alice", "alice@example.com", [])

    system.add_course(course)
    system.add_student(student)
    assert system.register_course(student, course) is True
    assert len(system.get_registered_courses(student)) == 1
