package main

import (
	"fmt"
	"std/github.com/bank/accounts"
)

func main() {
	account := accounts.NewAccouont("lsh")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
}
