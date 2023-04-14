package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}

	fmt.Println("Waiting for messages")
	resultOne := <-c
	resultTwo := <-c
	fmt.Println("Received this message:", resultOne)
	fmt.Println("Received this message:", resultTwo)

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	fmt.Println(person)
	c <- person + " is sexy"
}
