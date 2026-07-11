package digitalwallet

import (
	"math/big"
	"testing"
)

func TestDigitalWalletTransferFunds(t *testing.T) {
	t.Parallel()

	wallet := NewDigitalWallet()
	user1 := NewUser("U1", "Alice", "alice@example.com", "pass")
	user2 := NewUser("U2", "Bob", "bob@example.com", "pass")
	wallet.CreateUser(user1)
	wallet.CreateUser(user2)

	account1 := NewAccount("A1", user1, "111", USD)
	account2 := NewAccount("A2", user2, "222", USD)
	wallet.CreateAccount(account1)
	wallet.CreateAccount(account2)

	account1.Deposit(big.NewFloat(500))
	amount := big.NewFloat(100)

	if err := wallet.TransferFunds(account1, account2, amount, USD); err != nil {
		t.Fatalf("TransferFunds: %v", err)
	}

	balance1, _ := account1.GetBalance().Float64()
	balance2, _ := account2.GetBalance().Float64()
	if balance1 != 400 || balance2 != 100 {
		t.Fatalf("balances after transfer = (%v, %v), want (400, 100)", balance1, balance2)
	}
}
