package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return w.ErrInsufficientFunds(amount)
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) ErrInsufficientFunds(amount Bitcoin) error {
	errorMsg := fmt.Sprintf("withdrawal amount %s greater than current balance", amount.String())
	return errors.New(errorMsg)
}
