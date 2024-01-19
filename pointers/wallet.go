package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
  String() string
}

type Wallet struct {
  bitcoin Bitcoin
}

var ErrInsufficientFunds = errors.New("Not enough balance")

func (w *Wallet) Deposit(amount Bitcoin) {
  w.bitcoin += amount 
}

func (w *Wallet) Balance() Bitcoin {
  return (*w).bitcoin
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
  if amount > w.bitcoin {
    return ErrInsufficientFunds
  }
  w.bitcoin -= amount
  return nil
}

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}
