package linkedin

import "testing"

func TestLinkedInServiceRegisterAndLogin(t *testing.T) {
	t.Parallel()

	service := NewLinkedInService()
	user := NewUser("1", "Alice", "alice@example.com", "secret")
	service.RegisterUser(user)

	loggedIn, err := service.LoginUser("alice@example.com", "secret")
	if err != nil {
		t.Fatalf("LoginUser: %v", err)
	}
	if loggedIn.Name != "Alice" {
		t.Fatalf("logged in user = %q, want Alice", loggedIn.Name)
	}
}

func TestLinkedInServiceSearchUsers(t *testing.T) {
	t.Parallel()

	service := NewLinkedInService()
	service.RegisterUser(NewUser("1", "Alice", "alice@example.com", "pass"))
	service.RegisterUser(NewUser("2", "Bob", "bob@example.com", "pass"))

	results := service.SearchUsers("alice")
	if len(results) != 1 || results[0].Name != "Alice" {
		t.Fatalf("SearchUsers() = %+v, want one Alice match", results)
	}
}
