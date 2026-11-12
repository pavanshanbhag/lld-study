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


def test_search_courses_by_code() -> None:
    system = CourseRegistrationSystem()
    system.add_course(Course("CS101", "Intro", "Doe", 10, 0))
    system.add_course(Course("MATH200", "Calc", "Smith", 10, 0))
    results = system.search_courses("CS")
    assert len(results) == 1
    assert results[0].get_code() == "CS101"
