package main

import (
	"Project1/main.go/banking"
	"fmt"
)

func main() {
	account := banking.Account{Owner: "nicolas", Balance: 1000}
	fmt.Println(account)
}
