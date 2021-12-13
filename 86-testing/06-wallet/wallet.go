package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin float64

var ErrInsufficientFunds = errors.New("connot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.Balance() < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
