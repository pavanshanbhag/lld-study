package splitwise

import "testing"

func TestSplitwiseServiceAddUserAndGroup(t *testing.T) {
	t.Parallel()

	service := NewSplitwiseService()
	user := NewUser("1", "Alice", "alice@example.com")
	service.AddUser(user)

	group := NewGroup("1", "Trip")
	group.AddMember(user)
	service.AddGroup(group)

	if service.users["1"] == nil {
		t.Fatal("expected user to be registered")
	}
	if service.groups["1"] == nil {
		t.Fatal("expected group to be registered")
	}
}

func TestSplitwiseServiceAddExpense(t *testing.T) {
	t.Parallel()

	service := NewSplitwiseService()
	alice := NewUser("1", "Alice", "alice@example.com")
	bob := NewUser("2", "Bob", "bob@example.com")
	service.AddUser(alice)
	service.AddUser(bob)

	group := NewGroup("1", "Dinner")
	group.AddMember(alice)
	group.AddMember(bob)
	service.AddGroup(group)

	expense := NewExpense("1", 100, "Dinner", alice)
	expense.AddSplit(NewEqualSplit(alice))
	expense.AddSplit(NewEqualSplit(bob))
	service.AddExpense(group.ID, expense)

	key := alice.ID + ":" + bob.ID
	if alice.Balances[key] != 50 {
		t.Fatalf("alice balance with bob = %v, want 50", alice.Balances[key])
	}
}
