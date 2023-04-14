package main

import (
	"banking/Go-Nomad-Coder/Project2/2.4/mydict"
	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	fmt.Println(word)
}
