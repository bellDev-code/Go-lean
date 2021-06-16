package main

import (
	"fmt"

	"github.com/bellDev-code/leango/accounts"
)

func main() {
	account := accounts.NewAccount("jongho")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(20)
	fmt.Println(account.Balance())
}