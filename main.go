package main

import (
	"fmt"

	"github.com/bellDev-code/leango/accounts"
)

func main() {
	account := accounts.NewAccount("jongho")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}