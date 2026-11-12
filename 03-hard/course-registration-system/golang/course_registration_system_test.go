package courseregistrationsystem

import "testing"

func TestSmoke(t *testing.T) {
	t.Parallel()
	_ = NewCourseRegistrationSystem()
}
