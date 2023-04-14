package main

import (
	"banking/Go-Nomad-Coder/Project2/2.4/mydict"
	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	err := dictionary.Update("fsafa", "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)
}
