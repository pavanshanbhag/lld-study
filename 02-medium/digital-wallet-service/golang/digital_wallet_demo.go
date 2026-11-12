package digitalwallet

import (
	"log/slog"
	"math/big"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	digitalWallet := NewDigitalWallet()

	user1 := NewUser("U001", "John Doe", "john@example.com", "password123")
	user2 := NewUser("U002", "Jane Smith", "jane@example.com", "password456")
	digitalWallet.CreateUser(user1)
	digitalWallet.CreateUser(user2)

	account1 := NewAccount("A001", user1, "1234567890", USD)
	account2 := NewAccount("A002", user2, "9876543210", EUR)
	digitalWallet.CreateAccount(account1)
	digitalWallet.CreateAccount(account2)

	digitalWallet.AddPaymentMethod(NewCreditCard("PM001", user1, "1234567890123456", "12/25", "123"))
	digitalWallet.AddPaymentMethod(NewBankAccount("PM002", user2, "9876543210", "987654321"))

	account1.Deposit(big.NewFloat(1000.00))
	account2.Deposit(big.NewFloat(500.00))

	if err := digitalWallet.TransferFunds(account1, account2, big.NewFloat(100.00), USD); err != nil {
		logger.Error("transfer failed", "err", err)
	}

	for _, tx := range digitalWallet.GetTransactionHistory(account1) {
		logger.Info("account1 transaction", "id", tx.ID, "amount", tx.Amount, "currency", tx.Currency)
	}
	for _, tx := range digitalWallet.GetTransactionHistory(account2) {
		logger.Info("account2 transaction", "id", tx.ID, "amount", tx.Amount, "currency", tx.Currency)
	}
}
