package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(depositAmt Bitcoin) {
	w.balance += depositAmt
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(withdrawAmt Bitcoin) error {
	if withdrawAmt > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= withdrawAmt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
