package main

import (
	"fmt"

	"github.com/jeongola/Go-Nomad-Coder/Project1/banking"
)

func main() {
	account := banking.Account{Owner: "nicolas", Balance: 1000}
	fmt.Println(account)
}
