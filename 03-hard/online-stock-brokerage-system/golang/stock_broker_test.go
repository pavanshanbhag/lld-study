package onlinestockbrokeragesystem

import "testing"

func TestStockBrokerCreateAccountAndAddStock(t *testing.T) {
	t.Parallel()

	broker := NewStockBroker()
	user := NewUser("U1", "Alice", "alice@example.com")
	broker.CreateAccount(user, 10000)

	account := broker.GetAccount("A001")
	if account == nil {
		t.Fatal("expected account A001")
	}
	if account.GetBalance() != 10000 {
		t.Fatalf("balance = %v, want 10000", account.GetBalance())
	}

	apple := NewStock("AAPL", "Apple Inc.", 150)
	broker.AddStock(apple)

	if broker.GetStock("AAPL") == nil {
		t.Fatal("expected AAPL stock to be registered")
	}
}
