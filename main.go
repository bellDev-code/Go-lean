package main

import (
	"fmt"

	"github.com/bellDev-code/leango/accounts"
)

func main() {
	account := accounts.NewAccount("jongho")
	fmt.Println(account)
}