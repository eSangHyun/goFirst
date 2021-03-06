package accounts

import "errors"

// Account
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("no money")

// NewAccount create Account
func NewAccouont(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
