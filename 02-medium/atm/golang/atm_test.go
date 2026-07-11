package atm

import (
	"errors"
	"testing"
)

func TestATMWithdrawAndDeposit(t *testing.T) {
	t.Parallel()

	banking := NewBankingService()
	banking.CreateAccount("1234567890", 1000)

	dispenser := NewCashDispenser(5000)
	machine := NewATM(banking, dispenser)

	balance, err := machine.CheckBalance("1234567890")
	if err != nil || balance != 1000 {
		t.Fatalf("CheckBalance() = (%v, %v), want (1000, nil)", balance, err)
	}

	if err := machine.WithdrawCash("1234567890", 200); err != nil {
		t.Fatalf("WithdrawCash: %v", err)
	}

	balance, err = machine.CheckBalance("1234567890")
	if err != nil || balance != 800 {
		t.Fatalf("balance after withdraw = (%v, %v), want (800, nil)", balance, err)
	}

	if err := machine.DepositCash("1234567890", 50); err != nil {
		t.Fatalf("DepositCash: %v", err)
	}

	balance, err = machine.CheckBalance("1234567890")
	if err != nil || balance != 850 {
		t.Fatalf("balance after deposit = (%v, %v), want (850, nil)", balance, err)
	}
}

func TestATMInsufficientFunds(t *testing.T) {
	t.Parallel()

	banking := NewBankingService()
	banking.CreateAccount("1234567890", 100)

	machine := NewATM(banking, NewCashDispenser(5000))

	err := machine.WithdrawCash("1234567890", 200)
	if !errors.Is(err, ErrInsufficientFunds) {
		t.Fatalf("WithdrawCash err = %v, want ErrInsufficientFunds", err)
	}
}

func TestATMInsufficientCashInATM(t *testing.T) {
	t.Parallel()

	banking := NewBankingService()
	banking.CreateAccount("1234567890", 1000)

	machine := NewATM(banking, NewCashDispenser(100))

	err := machine.WithdrawCash("1234567890", 200)
	if !errors.Is(err, ErrInsufficientCashInATM) {
		t.Fatalf("WithdrawCash err = %v, want ErrInsufficientCashInATM", err)
	}

	balance, err := machine.CheckBalance("1234567890")
	if err != nil || balance != 1000 {
		t.Fatalf("balance should be unchanged: (%v, %v)", balance, err)
	}
}

func TestATMAccountNotFound(t *testing.T) {
	t.Parallel()

	machine := NewATM(NewBankingService(), NewCashDispenser(1000))

	_, err := machine.CheckBalance("missing")
	if err == nil {
		t.Fatal("expected error for missing account")
	}
}
