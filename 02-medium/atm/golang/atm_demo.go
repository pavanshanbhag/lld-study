package atm

import (
	"log/slog"
	"os"
)

func Run() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	bankingService := NewBankingService()
	cashDispenser := NewCashDispenser(10000)
	atmMachine := NewATM(bankingService, cashDispenser)

	bankingService.CreateAccount("1234567890", 1000.0)
	bankingService.CreateAccount("9876543210", 500.0)

	card := NewCard("1234567890", "1234")
	if err := atmMachine.AuthenticateUser(card); err != nil {
		logger.Error("authentication failed", "err", err)
		return
	}

	if balance, err := atmMachine.CheckBalance("1234567890"); err != nil {
		logger.Error("check balance failed", "err", err)
	} else {
		logger.Info("account balance", "balance", balance)
	}

	if err := atmMachine.WithdrawCash("1234567890", 500.0); err != nil {
		logger.Error("withdraw failed", "err", err)
	} else {
		logger.Info("withdrew cash", "amount", 500.0)
	}

	if err := atmMachine.DepositCash("9876543210", 200.0); err != nil {
		logger.Error("deposit failed", "err", err)
	} else {
		logger.Info("deposited cash", "amount", 200.0)
	}

	if balance, err := atmMachine.CheckBalance("1234567890"); err != nil {
		logger.Error("check balance failed", "err", err)
	} else {
		logger.Info("updated balance", "balance", balance)
	}
}
