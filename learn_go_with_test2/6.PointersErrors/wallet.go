package main

import (
    "errors"
    "fmt"
)

var ErrInsufficientBalance = errors.New("insufficient balance")

type Bitcoin int

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet)Deposit(m Bitcoin) {
    w.balance += m
}

func (w *Wallet)Balance() Bitcoin {
    return w.balance
}

func (w *Wallet)Withdraw(m Bitcoin) error {
    if m > w.balance {
        return ErrInsufficientBalance
    }
    w.balance -= m
    return nil
}

type Stringer interface {
    String() string
}

func (b Bitcoin)String() string {
    return fmt.Sprintf("%d BTC", b)
}



func main() {
    w := Wallet{Bitcoin(100)}

    w.Deposit(Bitcoin(100))
    fmt.Println(w.Balance())
}
