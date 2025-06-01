package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds error = errors.New("insufficient funds")

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of bitcoin in Deposit is %p \n", &w.balance)

	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
